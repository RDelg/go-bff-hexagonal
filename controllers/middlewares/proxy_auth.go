package middlewares

import (
	"bff/domain"
	"bff/domain/apperrors"
	"fmt"

	"github.com/gin-gonic/gin"
)

// ApigeeMiddleware creates an Apigee auth middleware
func ApigeeMiddleware(s domain.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := s.GetAccessToken()
		if err != nil {
			err := apperrors.NewAuthorization("Internal auth error")
			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			c.Abort()
			return
		}
		environment, _ := s.GetEnvironment()

		c.Request.Header.Set("authorization", fmt.Sprintf("Bearer %s", accessToken))
		c.Request.Header.Set("x-environment", environment)
		c.Next()

	}
}
