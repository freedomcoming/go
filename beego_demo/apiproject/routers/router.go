package routers

import (
	"apiproject/controllers"
	"github.com/astaxie/beego/plugins/cors"

	"github.com/astaxie/beego"
)

func init() {

	// 允许跨域请求
	beego.InsertFilter("*", beego.BeforeStatic, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.SetStaticPath("/static", "static")
	//ns := beego.NewNamespace("/v1",
	//	beego.NSNamespace("/object",
	//		beego.NSInclude(
	//			&controllers.ObjectController{},
	//		),
	//	),
	//	beego.NSNamespace("/user",
	//		beego.NSRouter("/user", &controllers.UserController{}),
	//	),
	//)
	//beego.AddNamespace(ns)
	beego.Router("/getall", &controllers.UserController{}, "*:GetAll")
	beego.Router("/get", &controllers.UserController{}, "*:Get")
	beego.Router("/login", &controllers.UserController{}, "*:Login")
	beego.Router("/post", &controllers.UserController{}, "post:Post")

}
