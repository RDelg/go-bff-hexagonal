package domain

import "github.com/gin-gonic/gin"

// ApigeePort contains all the methods to interactuate with ApiGee
type ApigeePort interface {
	GetAccessToken() (*ApigeeTokenClaims, error)
}

// AuthService defines methods the controller layer expects to interact
// with in regards to authorizing access
type AuthService interface {
	GetAccessToken() (string, error)
}

// HTTPService defines methods the controller layer expects to interact
// with in regards to sending and receiving http messages
type HTTPService interface {
	Get(string) gin.HandlerFunc
}
