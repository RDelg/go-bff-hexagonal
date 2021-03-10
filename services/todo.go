package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//FetchAllTodo fetches all
func FetchAllTodo(c *gin.Context) {
	fmt.Println("hello FetchAllTodo")
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "test": "123"})
}
