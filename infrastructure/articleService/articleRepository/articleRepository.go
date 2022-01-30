// DBから取得し、結果からエンティティ / 値オブジェクを生成して返却
package infrastructure

import (
	articleService "keep-back-api/infrastructure/articleService"
	infra "keep-back-api/infrastructure/db"
)

type ArticleRepository struct {
	new articleService.Article
}

var (
	dbService infra.DBService
	db        = dbService.DBSession()
	article   articleService.Article
)

func (r *ArticleRepository) newArticle() error {
	res, err := db.NamedExec("INSERT INTO articles (user_id, title, url, description, thumbnail, header, body) VALUES (:user_id, :title, :url, :description, :thumbnail, :header, :body)",
		map[string]interface{}{
			"user_id":     r.new.Userid,
			"title":       r.new.Title,
			"url":         r.new.URL,
			"description": r.new.Description,
			"thumbnail":   r.new.ThumbnailURL,
			"header":      r.new.Header,
			"body":        r.new.Body,
		})
	if err != nil {
		return err
	}
	r.new.ID, err = res.LastInsertId()
	return nil
}
