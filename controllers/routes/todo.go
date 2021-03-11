package routes

import (
	"bff/domain"

	"github.com/gin-gonic/gin"
)

// AddTodoRoutes adds todo routes
func AddTodoRoutes(config *domain.RestConfig, rg *gin.RouterGroup) {
	ping := rg.Group("/todo")
	ping.GET("/", config.HTTP.Get())
}
