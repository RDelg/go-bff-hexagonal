package routes

import (
	"bff/services"

	"github.com/gin-gonic/gin"
)

// AddTodoRoutes adds todo routes
func AddTodoRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/todo")

	ping.GET("/", services.FetchAllTodo)

}
