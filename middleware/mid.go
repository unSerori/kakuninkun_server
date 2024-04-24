package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ロギング
func MidLog(c *gin.Context) {
	fmt.Println("test log")
}
