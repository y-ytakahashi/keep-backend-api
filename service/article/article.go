package article

import (
		"fmt"
		// "encoding/json"
    "github.com/gin-gonic/gin"
    "readinglist-backend-api/db"
    "readinglist-backend-api/entity"
)

// Service procides Article's behavior
type Service struct{}

// Article is alias of entity.Article struct
type Article entity.Article
//APIリクエストの構造
type request struct{
	Userid 		string	`json:"user_id"`
	AuthToken string	`json:"authtoken"`
	URL 			string	`json:"url"`
}

// GetAll is get all Article
func (s Service) GetAll() ([]Article, error) {
    db := db.GetDB()
    var u []Article

    if err := db.Find(&u).Error; err != nil {
        return nil, err
    }

    return u, nil
}

// RequestParse front Request url param analysis OGP
func (s Service) RequestParse(c *gin.Context) (*gin.Context, error) {
	var u Article
	var r request

	// Frontからのリクエストパラメータのパースに失敗した場合に400を応答
	if err := c.BindJSON(&r); err == nil {
		fmt.Println("parse request param")
		c.AbortWithStatus(400)
		fmt.Println(err)
	}
  dummy := []byte(`{"num":6.13,"strs":["a","b"]}`)
	fmt.Println("do request param parse generate OGP")
	fmt.Println(dummy)
	// 構造体  BindJSON を呼び出す際に、独自データを代入する方法を調査

	u.Userid = r.Userid
	u.Title = `{"Title":"なぜぱーすされない"}`
	u.URL = r.URL
	u.Description = "description"
	u.Header = "header"
	u.Body = "Body"
	fmt.Println(u)

	return u, nil
}

// CreateModel is create Article model
// func (s Service) CreateModel(c *gin.Context) (Article, error) {
	func (s Service) CreateModel(c *gin.Context) (Article, error) {
    db := db.GetDB()
		// DB 登録用の構造
    var u Article

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