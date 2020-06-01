package model

import "time"

//ArticleInfo  文章结构体(不包含content)
type ArticleInfo struct {
	ID           int64     `db:"id"`
	CategoryID   int64     `db:"category_id"`
	Summary      string    `db:"summary"`
	Title        string    `db:"title"`
	ViewCount    uint32    `db:"view_count"`
	CommentCount uint32    `db:"comment_count"`
	UserName     string    `db:"username"`
	CreateTime   time.Time `db:"create_time"`
}

//ArticleDetail 用于文章详情页的实体
type ArticleDetail struct {
	ArticleInfo
	Category        //分类实体；用于相关内容推送
	Content  string `db:"content"`
}

//ArticleRecord 文章分页实体
type ArticleRecord struct {
	ArticleInfo
	Category
}

//ArticleD  文章实体用于接受文章表查询数据
type ArticleD struct{
	ArticleInfo
	Content  string `db:"content"`
}

//Search  用于搜索
type Search struct {
	Key string `form:"searchkey"`
}