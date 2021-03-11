package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPService contains the methods to acess apigee
type HTTPService struct{}

//Get gets data
func (s *HTTPService) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		acessToken, _ := c.Get("access_token")
		log.Printf("Hello from http service. acess_token: %v\n", fmt.Sprint(acessToken))
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "test": "123"})
	}
}
