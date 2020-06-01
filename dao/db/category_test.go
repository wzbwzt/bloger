package db

import (
	"testing"
)

func init() {
	//parseTime=true 将目前SQL中的时间类型自动解析为go结构体中的时间类型  不加报错
	dns := `root:123123@tcp(192.168.241.129:3306)/blogger?parseTime=true`
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}
func TestGetCategoryByID(t *testing.T) {
	category, err := GetCategoryByID(1)
	if err != nil {
		panic(err)
	}
	t.Logf("category:%#v", category)
}

func TestGetCategoryList(t *testing.T) {
	var categorylistID []int64
	categorylistID = append(categorylistID, 1, 2, 3)
	categoryList, err := GetCategoryList(categorylistID)
	if err != nil {
		panic(err)
	}
	for _, v := range categoryList {
		t.Logf("id:%d;category:%v\n", v.CategoryID, v)
	}

}

func TestGetAllCategory(t *testing.T) {
	clist, err := GetAllCategoryList()
	if err != nil {
		panic(err)
	}
	for _, v := range clist {
		t.Logf("id:%d,category:%v\n", v.CategoryID, v)
	}
}
