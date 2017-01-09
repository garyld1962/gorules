package main

import (
	R "gorules"
	"net/http"

	"fmt"

	"gopkg.in/gin-gonic/gin.v1"
)

// RuleMessage - structure of data sent to rules api.
type RuleMessage struct {
	Rule   interface{}            `form:"rule" json:"rule" binding:"required"`
	Action interface{}            `form:"action" json:"action"`
	Data   map[string]interface{} `form:"data" json:"data" binding:"required"`
}

// ExpressionMessage - structure for evaluating online expressions
type ExpressionMessage struct {
	Expression interface{}            `form:"expression" json:"expression" binding:"required"`
	Data       map[string]interface{} `form:"data" json:"data"`
}

func main() {
	router := gin.Default()

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "Service running"})
	})

	router.GET("/operators/math", func(c *gin.Context) {
		operators := R.MathOperatorList()
		c.JSON(http.StatusOK, gin.H{"operators": operators})
	})

	router.GET("/operators/string", func(c *gin.Context) {
		operators := R.StringOperatorList()
		c.JSON(http.StatusOK, gin.H{"operators": operators})
	})

	router.POST("/rule", func(c *gin.Context) {
		var json RuleMessage
		if c.BindJSON(&json) == nil {
			result := R.NewRuleProcessor(json.Rule).Process(json.Data)
			c.JSON(http.StatusOK, gin.H{"status": result})
		}

		// c.JSON(http.StatusOK, gin.H{"status": "Invalid data"})
	})

	router.POST("/expression", func(c *gin.Context) {
		var json ExpressionMessage
		if c.BindJSON(&json) == nil {
			result, _ := R.NewValue(json.Expression.(string)).Evaluate(json.Data)
			fmt.Println("json", json.Expression, result)
			c.JSON(http.StatusOK, gin.H{"result": result})
		}

		// c.JSON(http.StatusOK, gin.H{"status": "Invalid data"})
	})

	router.Run(":8080")
}
