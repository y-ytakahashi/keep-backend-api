package controllers

import (
	"fmt"
	usecase "keep-back-api/usecase/article"

	"github.com/gin-gonic/gin"
	// debug
	// "encoding/json"
)

// ArticleController is common controller
type ArticlePresentation struct{}

var articleUsecase usecase.ArticleService

// Create action: POST /article
func (s ArticlePresentation) New(c *gin.Context) {
	// r, err := service.RequestParse(c)
	p, err := articleUsecase.NewArticle(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

func (s ArticlePresentation) All(c *gin.Context) {
}
