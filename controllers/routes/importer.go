package routes

import (
	"bff/domain"
	"path"

	"github.com/gin-gonic/gin"
)

// AddRoutes adds all the routes
func AddRoutes(config *domain.Config, rg *gin.RouterGroup) {
	// TODO: Load urls from file and iterate
	url := "/test"
	rg.GET(url, config.HTTPService.Get(path.Join(config.APIEndpoint, url)))
}
