package models

import (
	"readinglist-backend-api/db"
	"readinglist-backend-api/entity"
)

// ArticleModels is tier model
type ArticleModels struct{}


// GetByID is get a Article
func (s ArticleModels) GetByID(id string) ([]entity.Article, error) {

	db := db.GetDB()
	var articles []entity.Article


	if err := db.Select("ID","Userid","ThumbnailURL","URL","Title","Description","CreatedAt").Where("Userid = ?",id).Find(&articles).Error; err != nil {
			return articles, err
	}


	return articles, nil
}

// GetViewReadingList フロントでリーディングリスト一覧を表示させるために使う
func (s ArticleModels) GetViewReadingList(id string) ([]entity.ViewArticle, error) {
	
	db := db.GetDB()
	var view []entity.ViewArticle
	var Article entity.Article

	if err := db.Model(&Article).Find(&view).Error; err != nil {
		return view, err

	}
	return view, nil
}
