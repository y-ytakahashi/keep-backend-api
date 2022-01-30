package domain

import (
	Article "keep-back-api/infrastructure/articleService/articleRepository"
	"time"
)

// ResArticle return front data
// type ArticleInfo struct {
// 	ID          uint   `gorm:"primarykey" gorm:"AUTO_INCREMENT"`
// 	Userid      string `json:"user_id"`
// 	Title       string `json:"title"`
// 	URL         string `json:"url"`
// 	Description string `json:"description"`
// 	Header      string `json:"header"`
// 	Body        string `json:"body"`
// }

type ArticleInfo struct {
	UserId       string
	Title        string
	Url          string
	Description  string
	Header       string
	ThumbnailURL string
	Body         string
	registedDate time.Time
}

var ArticleRepository Article.ArticleRepository

// func UserArticleLists(userID string) ([]ArticleInfo, error) {
// 	articles, err := ArticleRepository.GetByID(userID)
// 	return articles, err
// }

func RegisterArticle(article ArticleInfo) {
}
