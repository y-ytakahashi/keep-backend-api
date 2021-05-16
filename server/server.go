package server

import (
	"log"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"readinglist-backend-api/controller/user"
	"readinglist-backend-api/controller/article"
)

// Init is initialize server
func Init() {
    r := router()
    r.Run()
}

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

		// ユーザー関係のルーティング定義
    u := r.Group("/users")
    {
        ctrl := user.Controller{}
        u.GET("", ctrl.Index)
        u.GET("/:id", ctrl.Show)
        u.POST("", ctrl.Create)
        u.PUT("/:id", ctrl.Update)
        u.DELETE("/:id", ctrl.Delete)
    }

		a := r.Group("/article")
		{
			ctrl := article.Controller{}
			a.GET("",  ctrl.Index)
			a.POST("", ctrl.Create)
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

