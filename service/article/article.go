package article

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "readinglist-backend-api/db"
    "readinglist-backend-api/entity"


)

// Service procides Article's behavior
type Service struct{}

// Article is alias of entity.Article struct
type Article entity.Article

// ParseURL is Scan OGP Meta Data
type ParseURL ParseArticle
//APIリクエストの構造
type request struct{
	Userid 		string	`json:"user_id"`
	AuthToken   string	`json:"authtoken"`
    URL 		string	`json:"url"`
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
func (s Service) RequestParse(c *gin.Context) (Article, error) {
	var u Article
	var r request

	// Frontからのリクエストパラメータのパースに失敗した場合に400を応答
	if err := c.BindJSON(&r); err != nil {
		fmt.Println("parse request param")
		c.AbortWithStatus(400)
		fmt.Println(err)
	}

    // HTML 構造解析を実施
    res := ScanArticle(r.URL)
    fmt.Println(res)

	fmt.Println("do request param parse generate OGP")
	// ここに 対象Webページを解析するロジックを記述する
    //OGP から サムネイルなど概要用情報を取得する


    // 静的解析処理を記述するまで、OGPで取得できるデータはダミーとする
	u.Userid = r.Userid
	u.Title = res.Title
	u.URL = r.URL
	u.Description = res.Description
    u.ThumbnailURL = res.ThumbnailURL
	u.Header = res.Header
	u.Body = res.Body
    // if err := c.BindJSON(&u); err == nil {
	// 	fmt.Println("parse request param")
	// 	c.AbortWithStatus(400)
	// 	fmt.Println(err)
	// }
	return u, nil
}

// CreateModel is create Article model
func (s Service) CreateModel(c Article) (Article, error) {
    db := db.GetDB()
        // DB 登録用の構造
    // var u Article
    // err := json.Unmarshal([]byte(c), &u)
    // if err := c.BindJSON(&u); err != nil {
    // 	return u, err
    // }

    if err := db.Create(&c).Error; err != nil {
        return c, err
    }

    return c, nil
}


// GetByID is get a Article
func (s Service) GetByID(id string) ([]Article, error) {
    db := db.GetDB()
    var u []Article

    if err := db.Where("userid = ?", id).Find(&u).Error; err != nil {
        return u, err
    }

    return u, nil
}

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