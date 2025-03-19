package api

import "github.com/gin-gonic/gin"

func SuccessResponse() gin.H {
	return gin.H{
		"success": true,
	}
}
