package controllers

import (
	"bff/controllers/middlewares"
	"bff/controllers/routes"
	"bff/services"

	"github.com/gin-gonic/gin"
)

// RestConfig contains the Engine and
// the services
type RestConfig struct {
	R             *gin.Engine
	ApigeeService services.TokenService
}

// ConfigRouter creates a new rest server definition
func ConfigRouter(config *RestConfig) {
	v1 := config.R.Group("/v1")
	v1.Use(middlewares.ApigeeMiddleware(config.ApigeeService))
	routes.AddTodoRoutes(v1)
}
