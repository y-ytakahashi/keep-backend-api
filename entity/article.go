package entity

import (
	"gorm.io/gorm"
)


// Article is user models property
type Article struct {
		gorm.Model
		ID       			uint 			`gorm:"primarykey" gorm:"AUTO_INCREMENT"`
		Userid 				string 		`json:"user_id"`
    Title 				string 		`json:"title"`
		URL  					string 		`json:"url"`
		Description 	string 		`json:"description"`
		ThumbnailURL	string 		`json:"thumbnail"`
		Header 				string 		`json:"header"`
		Body 					string 		`json:"body"`
}

// ResArticle return front data
type ResArticle struct {
		ID        	uint 			`gorm:"primarykey" gorm:"AUTO_INCREMENT"`
		Userid 			string 		`json:"user_id"`
    Title 			string 		`json:"title"`
		URL  				string 		`json:"url"`
		Description string 		`json:"description"`
		Header 			string 		`json:"header"`
		Body 				string 		`json:"body"`
}

