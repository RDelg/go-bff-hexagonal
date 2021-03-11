package controllers

import (
	"bff/controllers/middlewares"
	"bff/controllers/routes"
	"bff/domain"
)

// ConfigRouter creates a new rest server definition
func ConfigRouter(config *domain.RestConfig) {
	v1 := config.R.Group("/v1")
	v1.Use(middlewares.ApigeeMiddleware(config.InternalAuthService))
	routes.AddTodoRoutes(config, v1)
}
