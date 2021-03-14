package domain

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProxyPort defines methods the controller layer expects to interact
// with in regards to sending and receiving http messages
type ProxyPort interface {
	Auth() (*TokenClaims, error)
	GetEnv() string
	DoRequest(string, string, *http.Header, io.Reader) (int, []byte, error)
}

// ProxyService defines methods the controller layer expects to interact
// with in regards to sending and receiving http messages
type ProxyService interface {
	AuthMiddleware() gin.HandlerFunc
	Do(string) gin.HandlerFunc
}
