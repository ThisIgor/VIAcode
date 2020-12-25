package base

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexAction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Welcome to Knowledge api v1"})
}

func ExamplePrivate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Private cotent"})
}
