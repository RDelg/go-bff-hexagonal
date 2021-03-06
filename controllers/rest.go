package controllers

import (
	"bff/controllers/routes"
	"bff/domain"
)

// ConfigRouter creates a new rest server definition
func ConfigRouter(config *domain.Config) {
	v1 := config.R.Group("/")
	v1.Use(config.ProxyService.AuthMiddleware())
	routes.AddRoutes(config, v1)
}
