package entity

import (
	"gorm.io/gorm"
)


// User is user models property
type Article struct {
		gorm.Model
		Userid 			string 		`json:"user_id"`
    Title 			string 		`json:"title"`
		URL  				string 		`json:"url"`
		Description string 		`json:"description"`
		Header 			string 		`json:"header"`
		Body 				string 		`json:"body"`
}