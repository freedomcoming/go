package main

import (
	"fmt"
	"gin_learn/middle"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type PostData struct {
	Name string `json:"name"`
	Age  int    `json:"age" binding:"required,mustang"`
	Age2 int    `json:"age2" binding:"required,mustang"`
}

// binding required 必填项
func mustang(f1 validator.FieldLevel) bool {
	fmt.Println(f1.Field().Interface().(int))
	return true

}

func MiddleFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("调用中间件前")
		c.Next()
		fmt.Println("调用中间件后")
	}

}

func main() {
	r := gin.Default()
	// 中间件加到路由分组里面 也叫请求拦截
	v1 := r.Group("v1").Use(MiddleFunc()).Use(middle.LogerMiddleware())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mustang", mustang)
	}

	v1.GET("ping/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := c.Query("name")
		age := c.DefaultQuery("age", "100")
		fmt.Println(id, user)
		c.JSON(200, gin.H{
			"message": "ok",
			"id":      id,
			"name":    user,
			"age":     age,
		})

	})
	r.POST("/test", func(c *gin.Context) {
		var p PostData
		err := c.ShouldBind(&p)
		fmt.Println(p.Name)
		fmt.Println(err)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 300,
				"data": p,
				"msg":  err,
			})
		} else {
			c.JSON(200, gin.H{
				"code": 200,
				"data": p,
			})
		}

	})

	r.POST("/upload", func(c *gin.Context) {
		//多文件
		//form := c.MultipartForm()
		//file := form.File["file"]
		file, _ := c.FormFile("file")                // 单文件
		c.SaveUploadedFile(file, "./"+file.Filename) // 保存文件
		c.JSON(200, gin.H{
			"msg": file,
		})
	})

	r.GET("middle", func(c *gin.Context) {
		fmt.Println(c.Request.Header.Get("User-Agent"))

		c.JSON(200, gin.H{
			"code": 200,
		})

	})

	// 中间件 路由访问前后都会走中间件
	r.Run()
}
