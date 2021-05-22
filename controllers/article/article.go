package controllers

import (
	"fmt"
	models "readinglist-backend-api/models/article"
	services "readinglist-backend-api/services/article"

	"github.com/gin-gonic/gin"
	// debug
	// "encoding/json"
)

// ArticleController is common controller
type ArticleController struct{}

var model models.ArticleModels
var service services.ArticleService
// Index action: GET /article
// Index action: GET /articles
// func (pc Controller) Index(c *gin.Context) {
// 	var s service.
// 	p, err := s.GetAll()

// 	if err != nil {
// 			c.AbortWithStatus(404)
// 			fmt.Println(err)
// 	} else {
// 			c.JSON(200, p)
// 	}
// }

// Show action: GET /users/:id
func (s ArticleController) Show(c *gin.Context) {
	// request が json の場合、 Query で値を取得する
	id := c.Param("id")
	fmt.Println("user id is", id)
	// p, err := service.GetByUserID(id)
	p, err := service.ViewArticleList(id)

	if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
	} else {
			c.JSON(200, p)
	}
}


// Create action: POST /article
func (s ArticleController) Create(c *gin.Context) {

	r, err := service.RequestParse(c)
	p, err := service.CreateModel(r)
	if err != nil {
			c.AbortWithStatus(400)
			fmt.Println(err)
	} else {
			c.JSON(201, p)
	}
}

