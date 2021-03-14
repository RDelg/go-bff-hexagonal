package services

import (
	"bff/domain"
	"bff/domain/apperrors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// ProxyService contains the methods to acess apigee
type ProxyService struct {
	Proxy domain.ProxyPort
}

// AuthMiddleware authenticates the service to the proxy
// and mutates the context header with the neccesary values
// to get authorized
func (s *ProxyService) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := s.Proxy.Auth()
		if err != nil {
			log.Printf("Error Auth: %v\n", err)
			err := apperrors.NewAuthorization("Proxy auth error")
			c.JSON(err.Status(), gin.H{"error": err})
			c.Abort()
			return
		}
		c.Request.Header.Set("authorization", fmt.Sprintf("Bearer %s", claims.AccessToken))
		c.Request.Header.Set("x-environment", s.Proxy.GetEnv())
		c.Next()
	}
}

// Do method
func (s *ProxyService) Do(method string) gin.HandlerFunc {
	return func(c *gin.Context) {
		status, body, err := s.Proxy.DoRequest(method, c.Request.URL.Path, &c.Request.Header, c.Request.Body)
		if err != nil {
			log.Printf("Error GET: %v\n", err)
			err := apperrors.NewBadRequest("Error sending the get request")
			c.JSON(err.Status(), gin.H{"error": err})
			c.Abort()
			return
		}
		c.String(status, string(body))
	}
}
