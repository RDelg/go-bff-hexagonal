package domain

import (
	"github.com/gin-gonic/gin"
)

// Auth defines methods the controller layer expects to interact
// with in regards to authorizing access
type Auth interface {
	GetAccessToken() (string, error)
}

// HTTP defines methods the controller layer expects to interact
// with in regards to sending and receiving http messages
type HTTP interface {
	Get(string) gin.HandlerFunc
}

// RestConfig contains the Engine and
// the services
type RestConfig struct {
	R                   *gin.Engine
	InternalAuthService Auth
	HTTP                HTTP
	APIEndpoint         string
}
