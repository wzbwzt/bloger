package main

import (
	"blog/controller"
	"blog/dao/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	dns := `root:123123@tcp(192.168.241.129:3306)/blogger?parseTime=true`
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	//加载静态文件
	r.Static("/relativepath", "./static")
	//加载资源图片
	r.Static("/imgs", "./tmp/imgs") 
	//加载模板
	r.LoadHTMLGlob("views/*")
	r.GET("/bloghome", controller.IndexHandler)
	r.POST("/bloghome", controller.SearchKey)
	r.GET("/category/",controller.CategoryList)
	r.GET("/article/",controller.GetArticleDetail)
	r.GET("/contribution",controller.PostContribution)
	r.POST("/contribution",controller.PostArticle)
	r.NoRoute(controller.Page404)
	r.Run(":8001")

}
