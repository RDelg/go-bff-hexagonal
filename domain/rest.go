package domain

import (
	"github.com/gin-gonic/gin"
)

// Config contains the Engine and
// the services
type Config struct {
	R                   *gin.Engine
	InternalAuthService AuthService
	ProxyService        ProxyService
	APIEndpoint         string
}
