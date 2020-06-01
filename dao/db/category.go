package db

import (
	"blog/model"

	"github.com/jmoiron/sqlx"
)

//InsertCategory 添加分类
func InsertCategory(c *model.Category) (categoryID int64, err error) {
	sqlStr := `insert inter category (category_name,category_no) value (?,?)`
	result, err := DB.Exec(sqlStr, c.CategoryName, c.CategoryNo)
	if err != nil {
		return
	}
	categoryID, err = result.LastInsertId()
	return
}

//GetCategoryByID 获取单个文章分类
func GetCategoryByID(id int64) (c *model.Category, err error) {
	c = &model.Category{}
	sqlStr := `select id,category_name,category_no from category where id=?`
	err = DB.Get(c, sqlStr, id)
	if err != nil {
		return
	}
	return
}

//GetCategoryList 获取多个查询分类
func GetCategoryList(categoryID []int64) (clist []*model.Category, err error) {
	sqlStr, args, err := sqlx.In("select id,category_name,category_no from category where id in (?)", categoryID) //构建sql语句
	if err != nil {
		return
	}
	err = DB.Select(&clist, sqlStr, args...)
	if err != nil {
		return
	}
	return
}

//GetAllCategoryList 获取全部分类
func GetAllCategoryList() (clist []*model.Category, err error) {
	sqlStr := `select id,category_name,category_no from category order by category_no asc`
	err = DB.Select(&clist, sqlStr)
	if err != nil {
		return
	}
	return
}
