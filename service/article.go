package service

import (
	"blog/dao/db"
	"blog/model"
)

//GetArticleRecordList 获取文章和对应的分类 (主要用于展示列表中既有文章同样也有分类信息)
func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	// 1.获取文章列表
	articleList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleList) <= 0 {
		return
	}
	//2.获取文章对应的分类
	ids := getCategoryIds(articleList)
	categoryList, err := db.GetCategoryList(ids)
	if err != nil {
		return
	}
	//返回页面 做聚合处理
	for _, articleInfo := range articleList {
		//根据当前文章来生成结构体
		articleRecordOne := &model.ArticleRecord{
			ArticleInfo: *articleInfo,
		}
		categoryid := articleInfo.CategoryID
		for _, category := range categoryList {
			if categoryid == category.CategoryID {
				articleRecordOne.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecordOne)
	}
	return
}

//根据文章列表来获取所有分类ids（去重）
func getCategoryIds(articleList []*model.ArticleInfo) (ids []int64) {
LABEL:
	for _, article := range articleList {
		categoryid := article.CategoryID
		for _, id := range ids {
			if categoryid == id {
				continue LABEL
			}
		}
		ids = append(ids, categoryid)
	}
	return
}

//GetArticleRecoldByid 跟据分类id获取该类文章和他对应的分类信息
func GetArticleRecoldByid(categoryID, pageNum, pageSize int) (articleRecoldList []*model.ArticleRecord, err error) {
	articleInfoList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}
	category, err := db.GetCategoryByID(int64(categoryID))
	if err != nil {
		return
	}
	if category == nil {
		return
	}
	for _, articleInfo := range articleInfoList {
		articleRecord := &model.ArticleRecord{
		}
		id := articleInfo.CategoryID
		if categoryID == int(id) {
			articleRecord.ArticleInfo=*articleInfo
			articleRecord.Category = *category
			articleRecoldList = append(articleRecoldList, articleRecord)
		}
	}
	return
}

// GetLatestArticleInfoList 获取最新文章
func GetLatestArticleInfoList(pageSize int)(articleList []*model.ArticleInfo,err error){
	articleList, err = db.GetArticleList(0, pageSize)
	if err!=nil{
		return
	}
	return
}

//GetArticleAllData 根据文章id ，获取文章详情，包含分类信息
func  GetArticleAllData(id int64)(articleDetail *model.ArticleDetail,err error){
	articleD, err := db.GetArticleByarticleid(id)
	if err != nil {
		return
	}
	categoryid:= articleD.ArticleInfo.CategoryID
	category, err := db.GetCategoryByID(categoryid)
	if err != nil {
		return
	}
	articleDetail=&model.ArticleDetail{
	}
	articleDetail.Category=*category
	articleDetail.ArticleInfo=articleD.ArticleInfo
	articleDetail.Content=articleD.Content
	return
}
//GetArticleBySearch 标题关键字搜索
func GetArticleBySearch(search *model.Search, pageNum,pageSize int)(articleRecordList []*model.ArticleRecord, err error){
	key:=(*search).Key
	articleInfoList, err := db.GetArticleByKey(key, pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}

	//2.获取文章对应的分类
	ids := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(ids)
	if err != nil {
		return
	}
	//返回页面 做聚合处理
	for _, articleInfo := range articleInfoList {
		//根据当前文章来生成结构体
		articleRecordOne := &model.ArticleRecord{
			ArticleInfo: *articleInfo,
		}
		categoryid := articleInfo.CategoryID
		for _, category := range categoryList {
			if categoryid == category.CategoryID {
				articleRecordOne.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecordOne)
	}
	return
}
