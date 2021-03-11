package middlewares

import (
	"bff/domain"
	"bff/domain/apperrors"

	"github.com/gin-gonic/gin"
)

// ApigeeMiddleware creates an Apigee auth middleware
func ApigeeMiddleware(s domain.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := s.GetAccessToken()
		if err != nil {
			err := apperrors.NewAuthorization("Provided token is invalid")
			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			c.Abort()
			return
		}
		c.Set("access_token", accessToken)
		c.Next()

	}
}
