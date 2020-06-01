package service

//业务逻辑层
import (
	"blog/dao/db"
	"blog/model"
)

//GetCategory 获取所有分类 （用于分类云）
func GetCategory() (categorylist []*model.Category, err error) {
	categorylist, err = db.GetAllCategoryList()
	if err != nil {
		return
	}
	return
}
