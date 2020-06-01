package controller

import (
	"blog/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//IndexHandler 访问主页的控制器
func IndexHandler(c *gin.Context) {
	//从service层获取数据
	//1.加载文章数据
	articleRecordList, err := service.GetArticleRecordList(0, 6)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//2.加载分类数据
	categoryList, err := service.GetCategory()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	articleLatestList,err := service.GetLatestArticleInfoList(3)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/blogHome.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
		"articleLatest_list":articleLatestList,
	})

}

//CategoryList 点击分类云
func CategoryList(c *gin.Context) {
	categoryidStr := c.Query("category_id") //从url链接中取值
	categoryid, err := strconv.ParseInt(categoryidStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	articleRecoldList, err := service.GetArticleRecoldByid(int(categoryid), 0, 6)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	categoryList, err := service.GetCategory()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	articleLatestList,err := service.GetLatestArticleInfoList(3)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/blogHome.html", gin.H{
		"article_list":  articleRecoldList,
		"category_list": categoryList,
		"articleLatest_list":articleLatestList,

	})
}

//404页面
func Page404(c *gin.Context){
	c.HTML(http.StatusNotFound,"views/404.html",nil)

}


