package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateParam(paramName string, rules ...func(s string) (ok bool, msg string)) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param(paramName)
		for rule := range rules {
			ok, msg := rules[rule](idStr)
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": msg})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
