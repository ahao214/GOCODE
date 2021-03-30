package service

import (
	"Ahao.Blogger/dao/db"
	"Ahao.Blogger/model"
	"fmt"
)

func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	articleInfoList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil {
		fmt.Printf("get aritcle list failed,err:%v\n", err)
		return
	}
	if len(articleInfoList) == 0 {
		return
	}
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		fmt.Printf("get category list failed,err:%v\n", err)
		return
	}
	for _, article := range articleInfoList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.CategoryId
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {

}
