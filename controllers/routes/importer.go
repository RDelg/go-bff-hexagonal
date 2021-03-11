package routes

import (
	"bff/domain"
	"path"

	"github.com/gin-gonic/gin"
)

// AddRoutes adds all the routes
func AddRoutes(config *domain.RestConfig, rg *gin.RouterGroup) {
	// TODO: Load urls from file and iterate
	url := "/test"
	rg.GET(url, config.HTTP.Get(path.Join(config.APIEndpoint, url)))
}
