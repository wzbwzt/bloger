package db

import (
	"blog/model"
)

// InsertArticle 插入文章
func InsertArticle(articleDetail *model.ArticleDetail) (articleID int64, err error) {
	if articleDetail == nil {
		return
	}
	sqlStr := `insert into article (category_id,summary,title,view_count,comment_count,username,content,create_time) values(?,?,?,?,?,?,?,?)`
	result, err := DB.Exec(sqlStr, articleDetail.ArticleInfo.CategoryID, articleDetail.ArticleInfo.Summary, articleDetail.ArticleInfo.Title,
		articleDetail.ArticleInfo.ViewCount, articleDetail.ArticleInfo.CommentCount, articleDetail.ArticleInfo.UserName, articleDetail.Content, articleDetail.ArticleInfo.CreateTime)
	if err != nil {
		return
	}
	articleID, err = result.LastInsertId()
	return
}

//GetArticleList 获取列表  设置文章分页
func GetArticleList(pageNum, pageSize int) (article []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize < 0 {
		return
	}
	sqlStr := `select id,category_id,summary,title,view_count,comment_count,username,create_time 
	from article where status=1 order by create_time desc limit ?,?`
	err = DB.Select(&article, sqlStr, pageNum, pageSize)
	return
}

//GetArticleByarticleid 查询单个文章（根据文章id）
func GetArticleByarticleid(articleID int64) (articleD *model.ArticleD, err error) {
	articleD=&model.ArticleD{}
	if articleID < 0 {
		return
	}
	sqlStr := `select id,category_id,summary,title,view_count,comment_count,username,create_time,content from article where status=1 and id=?`
	err = DB.Get(articleD, sqlStr, articleID)
	return
}

//GetArticleBycategoryid 查询一类文章（根据分类id）
func GetArticleBycategoryid(categoryID, pageNum, pageSize int64) (articleList []*model.ArticleInfo, err error) {
	if categoryID < 0 || pageNum < 0 || pageSize < 0 {
		return
	}
	sqlStr := `select id,category_id,summary,title,view_count,comment_count,username,create_time from article 
	where status and category_id=? order by create_time desc limit ?,?`
	err = DB.Select(&articleList, sqlStr, categoryID, pageNum, pageSize)
	return
}

//GetArticleByKey  根据文章标题来关键字搜索
func GetArticleByKey(key string ,pageNum int, pageSize int )(article []*model.ArticleInfo, err error){
	sqlStr:=`select id,category_id,summary,title,view_count,comment_count,username,create_time 
	from article where status=1 and title like "%?%" order by create_time desc limit ?,?`
	err = DB.Select(&article, sqlStr, key, pageNum, pageSize)
	return
}