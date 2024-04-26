package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// json
func TestJson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ // bodyがJSON形式のレスポンスを返す
		"message": "hello go server!",
	})
}
