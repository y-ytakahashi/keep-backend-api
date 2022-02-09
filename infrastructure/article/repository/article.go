// DBから取得し、結果からエンティティ / 値オブジェクを生成して返却
package infrastructure

import (
	"fmt"
	domain "keep-back-api/domain/model/article"
	infra "keep-back-api/infrastructure/db"
	"time"
)

type ArticleRepository struct{}

var dbService infra.DBService

func (ArticleRepository) AllArticle() []domain.ArticleInfo {
	fmt.Println("DBからすべてのArticleを取得")
	db := dbService.Con()

	var articles []domain.ArticleInfo
	// 取得したデータはページネーションできるようにしたい
	db.Select(&articles, "SELECT * FROM articles")

	return articles
}

func (ArticleRepository) NewArticle(article domain.ArticleInfo) error {
	fmt.Println("Article の情報をDBに保存")
	db := dbService.Con()
	db.DB.Ping()

	_, err := db.NamedExec(`
		INSERT INTO articles (user_id, title, url, description, thumbnail, header, body)
		VALUES (:user_id, :title, :url, :description, :thumbnail, :header, :body)
	`,
		map[string]interface{}{
			"user_id":     article.UserId,
			"title":       article.Title,
			"url":         article.Url,
			"description": article.Description,
			"thumbnail":   article.ThumbnailURL,
			"header":      article.Header,
			"body":        article.Body,
			"created_at":  time.Now(),
			"modified_at": time.Now(),
		})
	if err != nil {
		return err
	}

	return nil
}
