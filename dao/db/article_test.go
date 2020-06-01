package db

import (
	"blog/model"
	"testing"
	"time"
)

func init() {
	//parseTime=true 将目前SQL中的时间类型自动解析为go结构体中的时间类型  不加报错
	dns := `root:123123@tcp(192.168.241.129:3306)/blogger?parseTime=true`
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

//测试插入文章
func TestInsertArticle(t *testing.T) {
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryID = 1
	article.ArticleInfo.Summary = "摘要1"
	article.ArticleInfo.Title = "标题1"
	article.ArticleInfo.ViewCount = 1
	article.ArticleInfo.CommentCount = 1
	article.ArticleInfo.UserName = "wzb"
	article.Content = "内容1"
	article.ArticleInfo.CreateTime = time.Now()
	articleid, err := InsertArticle(article)
	if err != nil {
		t.Logf("error:%v", err)
	}
	t.Logf("articleID:%d", articleid)

}

//测试获取文章列表
func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(2, 2)
	if err != nil {
		t.Logf("error:%v", err)
	}
	t.Logf("articlelen:%d", len(articleList))
	for _, v := range articleList {
		t.Logf("articleID:%#v", v)
	}
}

//测试根据文章id获取详情
func TestGetArticleByarticleid(t *testing.T) {
	articleDetail, err := GetArticleByarticleid(2)
	if err != nil {
		t.Logf("error:%v", err)
	}
	t.Logf("detail:%v\n", articleDetail)
}
