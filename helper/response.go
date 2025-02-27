package helper

import "github.com/gin-gonic/gin"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseSuccess(message string, data interface{}) gin.H {
	return gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	}
}

func ResponseFailed(message string) gin.H {
	return gin.H{
		"status":  "failed",
		"message": message,
	}
}
