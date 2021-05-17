package article
import (
	"fmt"
	// "log"
	"github.com/gin-gonic/gin"
	"readinglist-backend-api/service/article"
	// debug
  // "encoding/json"
)

// Controller is article controller
type Controller struct{}

// Index action: GET /article
func (pc Controller) Index(c *gin.Context) {
	var s article.Service
	p, err := s.GetAll()

	if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
	} else {
			c.JSON(200, p)
	}
}

// Create action: POST /article
func (pc Controller) Create(c *gin.Context) {
	var s article.Service



	r, err := s.RequestParse(c)

	p, err := s.CreateModel(r)


	if err != nil {
			c.AbortWithStatus(400)
			fmt.Println(err)
	} else {
			c.JSON(201, p)
	}
}

