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
			err := apperrors.NewAuthorization("Internal auth error")
			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			c.Abort()
			return
		}
		c.Request.Header.Set("authorization", fmt.Sprintf("Bearer %s", claims.AccessToken))
		c.Request.Header.Set("x-environment", s.Proxy.GetEnv())
		c.Next()
	}
}

// Get method
func (s *ProxyService) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		status, body, err := s.Proxy.Get(c.Request.URL.Path, &c.Request.Header)
		if err != nil {
			log.Println(err)
		}
		c.String(status, string(body))
	}
}

// Post method
func (s *ProxyService) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		status, body, err := s.Proxy.Post(c.Request.URL.Path, &c.Request.Header, c.Request.Body)
		if err != nil {
			log.Println(err)
		}
		c.String(status, string(body))
	}
}
