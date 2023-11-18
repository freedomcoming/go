package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Hello struct {
	gorm.Model
	Name      string `gorm:"primary_key;column:nameNew;type:varchar(100);not null"sql:"index"`
	Sex       bool
	Age       int
	Hobbies   []Hobby    // 一对多
	IDCard    Card       // 一对一
	Languages []Language `gorm:"many2many:user_languages;"` // 多对多
}

type Hobby struct {
	gorm.Model

	Name    string
	HelloID int // 一对多
}

type Card struct {
	gorm.Model
	Name    string
	HelloID int // 一对一

}

type Language struct {
	gorm.Model
	Name string
}

func (h Hello) TableName() string {
	// 更改表名
	return "user"

}

func main() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3308)/ginclass?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	// 自动同步表 需要新建库
	db.AutoMigrate(&Hello{}, &Hobby{}, &Card{}, &Language{})
	// 创建记录
	c := Card{
		Name: "card",
	}
	l := Language{
		Name: "chinese",
	}
	ho := Hobby{
		Name: "fuck",
	}
	h := Hello{
		Name:      "wxw",
		Sex:       true,
		Age:       100,
		Hobbies:   []Hobby{ho},
		IDCard:    c,
		Languages: []Language{l},
	}
	db.Create(&h)

	db.Create(&ho)
	// 查询
	//var hello Hello
	//var hellos []Hello
	//
	//db.Where("nameNew=?", "wxw").Find(&hello)
	//fmt.Println(hello)
	//
	//db.Find(&hellos)
	//fmt.Println(hellos)
	//
	//for i, v := range hellos {
	//	fmt.Println(i, v)
	//}
	//
	//// update更新
	//
	//db.Where("name=?", "wxw").First(&Hello{}).Update("age", 10000)
	//
	//// delete操作 把删除时间戳加上 不是真正删除 加 .upscoped 硬删除
	//db.Where("name=?", "wxw").Delete(&Hello{})
	var h1 Hello
	//db.Where("id=?", 6).First(&h1)
	// 预加载模式
	db.Preload("IDCard").Preload("Languages").Where("id=?", 6).First(&h1) // 预加载
	fmt.Println(h1)

	defer db.Close()
}
