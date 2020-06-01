package controller

import (
	"blog/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PostContribution 加载贡献页
func PostContribution(c *gin.Context){
	categoryList, err := service.GetCategory()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK,"views/contribution.html",gin.H{
	"category_list":categoryList,
})
}

// PostArticle 文件提交
func PostArticle(c *gin.Context){
	//file, err := c.FormFile("f1")
	form,err:= c.MultipartForm()
	files := form.File["file"]
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	for index,file:= range files{
		// 上传文件到指定的目录
		dst:= fmt.Sprintf("D:/www/blog/tmp/files/(%d)%s",index,file.Filename)
		c.SaveUploadedFile(file,dst)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%d' file has uploaded!", len(files)),
		"data":form,
	})
}