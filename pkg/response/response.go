package response

import "github.com/gin-gonic/gin"

type envelope struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func Success(c *gin.Context, status int, data any) {
	c.JSON(status, envelope{Success: true, Data: data})
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, envelope{Success: false, Error: message})
}
