package main

import (
	R "gorules"
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

// Message - structure of data sent to rules api.
type Message struct {
	Rule   interface{}            `form:"rule" json:"rule" binding:"required"`
	Action interface{}            `form:"action" json:"action"`
	Data   map[string]interface{} `form:"data" json:"data" binding:"required"`
}

func main() {
	router := gin.Default()

	router.POST("/rules", func(c *gin.Context) {
		var json Message
		if c.BindJSON(&json) == nil {
			result := R.EvaluateRules(R.NewRuleFetcher(json.Rule.(string)), json.Data)
			c.JSON(http.StatusOK, gin.H{"status": result})
		}
	})

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "Service running"})
	})

	router.Run(":8080")
}
