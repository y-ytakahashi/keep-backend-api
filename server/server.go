package server

import (
	"log"
	"time"

	articleApp "keep-back-api/presentation/api.article"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct{}

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

// ArticleController is article
var articleRoute articleApp.ArticlePresentation

// TODO ルーティングの実装を 実践GOを参考にコントローラーとルーターの２役に分離する
func router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	a := r.Group("/articles")
	{
		a.POST("", articleRoute.New)
		// a.GET("/:id", ArticleController.Show)
	}

	return r
}

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before logic")
		// c.Next の前に記述した処理はルーティング内の処理の前に実行される
		c.Next()
		// c.Next の後に記述した処理はルーティング後の処理の後に実行される
		log.Println("after logic")
	}
}
