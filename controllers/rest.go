package controllers

import (
	"bff/controllers/middlewares"
	"bff/controllers/routes"
	"bff/domain"
)

// ConfigRouter creates a new rest server definition
func ConfigRouter(config *domain.Config) {
	v1 := config.R.Group("/")
	v1.Use(middlewares.ApigeeMiddleware(config.InternalAuthService))
	routes.AddRoutes(config, v1)
}
