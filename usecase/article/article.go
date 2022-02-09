package usecase

import (
	"fmt"
	domain "keep-back-api/domain/model/article"
	infra "keep-back-api/infrastructure/article/repository"

	"github.com/gin-gonic/gin"
)

// ArticleService procides Article's behavior
type ArticleService struct{}

// APIリクエストの構造
var repo infra.ArticleRepository

// RequestParse front Request url param analysis OGP
func (s ArticleService) NewArticle(c *gin.Context) (domain.ArticleInfo, error) {
	var ArticleInfo domain.ArticleInfo
	var clientRequest domain.ArticleClientRequest

	fmt.Println("do new article process")

	// Frontからのリクエストパラメータのパースに失敗した場合に400を応答
	if err := c.BindJSON(&clientRequest); err != nil {
		// fmt.Println("parse request param")
		c.AbortWithStatus(400)
	}

	// HTML 構造解析を実施
	ArticleInfo = domain.CreateArticleInfo(clientRequest)
	res := repo.NewArticle(ArticleInfo)

	fmt.Println("do request param parse generate OGP")
	return ArticleInfo, res
}
