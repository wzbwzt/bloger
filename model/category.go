package model

//Category 分类结构体
type Category struct {
	CategoryID   int64  `db:"id"`
	CategoryName string `db:"category_name"`
	CategoryNo   int    `db:"category_no"`
}
