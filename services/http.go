package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPService contains the methods to acess apigee
type HTTPService struct{}

//Get gets data
func (s *HTTPService) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		acessToken, _ := c.Get("access_token")
		fmt.Printf("hello Get %v\n", fmt.Sprint(acessToken))
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "test": "123"})
	}
}
