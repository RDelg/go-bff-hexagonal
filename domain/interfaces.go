package domain

import "github.com/gin-gonic/gin"

// ProxyAuthPort contains all the methods to interactuate with ApiGee
type ProxyAuthPort interface {
	GetAccessToken() (*ApigeeTokenClaims, error)
}

// AuthService defines methods the controller layer expects to interact
// with in regards to authorizing access
type AuthService interface {
	GetAccessToken() (string, error)
	GetEnvironment() (string, error)
}

// ProxyService defines methods the controller layer expects to interact
// with in regards to sending and receiving http messages
type ProxyService interface {
	Get() gin.HandlerFunc
	Post() gin.HandlerFunc
}
