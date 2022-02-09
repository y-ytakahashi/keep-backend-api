package infrastructure

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Article is user models property
type Article struct {
	ID           int64     `db:id`
	Userid       string    `db:"user_id"`
	Title        string    `db:"title"`
	URL          string    `db:"url"`
	Description  string    `db:"description"`
	ThumbnailURL string    `db:"thumbnail"`
	Header       string    `db:"header"`
	Body         string    `db:"body"`
	CreatedAt    time.Time `db:"created_at"`
	ModifiedAt   time.Time `db:"modified_at"`
}

// ViewArticle フロントに返却する想定のデータ構造
type ViewArticle struct {
	ID           int64  `db:id`
	Userid       string `db:"user_id"`
	Title        string `db:"title"`
	URL          string `db:"url"`
	Description  string `db:"description"`
	ThumbnailURL string `db:"thumbnail"`
}
