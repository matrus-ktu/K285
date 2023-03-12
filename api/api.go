package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestAPI(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "test",
	})
}

func Error(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "error",
	})
}
