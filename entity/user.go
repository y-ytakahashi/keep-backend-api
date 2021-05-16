package entity

import (
	"gorm.io/gorm"
)


// User is user models property
type User struct {
		gorm.Model
		ID        	uint 			`gorm:"primarykey" gorm:"AUTO_INCREMENT"`
		Email 			string		`json:"email"`
    FirstName 	string 		`json:"firstname"`
    LastName  	string 		`json:"lastname"`
		Password 		string 		`json:"password"`
}