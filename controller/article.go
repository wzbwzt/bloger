package controller

import (
	"blog/model"
	"blog/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//文章详情页
func GetArticleDetail(c *gin.Context ){
	idStr := c.Query("article_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	//根据文章id获取文章详细信息
	articleDetail, err := service.GetArticleAllData(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	//标签云列表
	categoryList, err := service.GetCategory()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//获取最新文章列表
	articleLatestList,err := service.GetLatestArticleInfoList(3)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK,"views/articleDetail.html",gin.H{
		"article_detail":articleDetail,
		"category_list":categoryList,
		"articleLatest_list":articleLatestList,
	})
}

//SearchKey  标题关键字查询
func SearchKey(c *gin.Context){
	var searchStruct model.Search
    err:=c.ShouldBind(&searchStruct)
	if err != nil {
		c.JSON(http.StatusNotFound,gin.H{"error":err})
		return
	}
	searchList, err := service.GetArticleBySearch(&searchStruct)

}
