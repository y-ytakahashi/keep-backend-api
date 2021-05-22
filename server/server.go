package server

import (
	"log"

	controllers "readinglist-backend-api/controllers/article"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {}


// Init is initialize server
func Init() {
    r := router()
    r.Run()
}
// ArticleController is article
var ArticleController controllers.ArticleController

func router() *gin.Engine {
    r := gin.Default()
		r.Use(middleware())
		r.Use(cors.New(cors.Config{
			AllowOrigins: []string{
        "http://localhost:3000",
    },
		AllowCredentials: true,
		MaxAge: 24 * time.Hour,
		}))

		// // ユーザー関係のルーティング定義
    // u := r.Group("/users")
    // {
		// 	ctrl := user.Controller{}
		// 	u.GET("", ctrl.Index)
		// 	u.GET("/:id", ctrl.Show)
		// 	u.POST("", ctrl.Create)
		// 	u.PUT("/:id", ctrl.Update)
		// 	u.DELETE("/:id", ctrl.Delete)
		// }

		a := r.Group("/articles")
		{

			a.POST("", ArticleController.Create)
			a.GET("/:id",  ArticleController.Show)
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

