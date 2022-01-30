package usecase

import (
	"fmt"
	Article "keep-back-api/domain/article"

	"github.com/gin-gonic/gin"
)

// ArticleService procides Article's behavior
type ArticleService struct{}

// APIリクエストの構造
type request struct {
	Userid    string `json:"user_id"`
	AuthToken string `json:"authtoken"`
	URL       string `json:"url"`
}

// RequestParse front Request url param analysis OGP
func (s ArticleService) NewArticle(c *gin.Context) (Article.ArticleInfo, error) {
	var ArticleInfo Article.ArticleInfo
	var clientRequest request

	// Frontからのリクエストパラメータのパースに失敗した場合に400を応答
	if err := c.BindJSON(&clientRequest); err != nil {
		// fmt.Println("parse request param")
		c.AbortWithStatus(400)
		// fmt.Println(err)
	}

	// HTML 構造解析を実施
	scanResult := ScanArticle(clientRequest.URL)
	fmt.Println(scanResult)

	fmt.Println("do request param parse generate OGP")
	// ここに 対象Webページを解析するロジックを記述する
	// OGP から サムネイルなど概要用情報を取得する

	// 静的解析処理を記述するまで、OGPで取得できるデータはダミーとする

	ArticleInfo.UserId = clientRequest.Userid
	ArticleInfo.Title = scanResult.Title
	ArticleInfo.Url = scanResult.URL
	ArticleInfo.Description = scanResult.Description
	ArticleInfo.ThumbnailURL = scanResult.ThumbnailURL
	ArticleInfo.Header = scanResult.Header
	ArticleInfo.Body = scanResult.Body
	// if err := c.BindJSON(&u); err == nil {
	// 	fmt.Println("parse request param")
	// 	c.AbortWithStatus(400)
	// 	fmt.Println(err)
	// }
	return ArticleInfo, nil
}

// // GetByUserID is get a Article
// func (s ArticleService) GetByUserID(id string) ([]entity.Article, error) {
// 	var methods models.ArticleModels
// 	fmt.Println(id)
// 	res, err := methods.GetByID(id)
// 	if err != nil {
// 		fmt.Println("error")
// 		return res, err
// 	}

// 	return res, nil
// }

// // GetByArticle
// func (s ArticleService) ViewArticleList(id string) ([]entity.ViewArticle, error) {
// 	var methods models.ArticleModels
// 	fmt.Println(id)
// 	res, err := methods.GetViewReadingList(id)
// 	if err != nil {
// 		fmt.Println("error")
// 		return res, err
// 	}

// 	return res, nil
// }
