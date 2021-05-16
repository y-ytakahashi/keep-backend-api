package article

import (
		"fmt"
		"time"
    "github.com/gin-gonic/gin"
    "readinglist-backend-api/db"
    "readinglist-backend-api/entity"
)

// Service procides Article's behavior
type Service struct{}

// Article is alias of entity.Article struct
type Article entity.Article

// GetAll is get all Article
func (s Service) GetAll() ([]Article, error) {
    db := db.GetDB()
    var u []Article

    if err := db.Find(&u).Error; err != nil {
        return nil, err
    }

    return u, nil
}

// CreateModel is create Article model
func (s Service) CreateModel(c *gin.Context) (Article, error) {
    db := db.GetDB()
    var u Article
		u.CreatedAt = time.Now()
		u.UpdatedAt = time.Now()
		fmt.Println(c.Request)

    if err := c.BindJSON(&u); err != nil {
        return u, err
    }

    if err := db.Create(&u).Error; err != nil {
        return u, err
    }

    return u, nil
}

// // GetByID is get a Article
// func (s Service) GetByID(id string) (Article, error) {
//     db := db.GetDB()
//     var u Article

//     if err := db.Where("id = ?", id).First(&u).Error; err != nil {
//         return u, err
//     }

//     return u, nil
// }

// // UpdateByID is update a Article
// func (s Service) UpdateByID(id string, c *gin.Context) (Article, error) {
//     db := db.GetDB()
//     var u Article

//     if err := db.Where("id = ?", id).First(&u).Error; err != nil {
//         return u, err
//     }

//     if err := c.BindJSON(&u); err != nil {
//         return u, err
//     }

//     db.Save(&u)

//     return u, nil
// }

// // DeleteByID is delete a Article
// func (s Service) DeleteByID(id string) error {
//     db := db.GetDB()
//     var u Article

//     if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
//         return err
//     }

//     return nil
// }