package entity

import (
	"gorm.io/gorm"
)

// User is user models property
type Leadinglist struct {

		ID		 			uint 		`gorm:"primarykey" gorm:"AUTO_INCREMENT"`
    ArticleID 	uint 		`json:"article_id"`
		ArticleMemo string 	`json:"article_memo"`
		IsPublish 	bool 		`json:"is_publish"`
		gorm.Model
}