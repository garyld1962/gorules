package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	R "gorules"

	"github.com/gin-gonic/gin"
)

// RuleMessage - structure of data sent to rules api.
type RuleMessage struct {
	Rule   interface{}            `form:"rule" json:"rule" binding:"required"`
	Action interface{}            `form:"action" json:"action"`
	Data   map[string]interface{} `form:"data" json:"data" binding:"required"`
}

// ExpressionMessage - structure for evaluating online expressions
type ExpressionMessage struct {
	Expression string                 `form:"expression" json:"expression" binding:"required"`
	Data       map[string]interface{} `form:"data" json:"data"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

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
		var msg RuleMessage
		if err := c.ShouldBindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := R.NewRuleProcessor(msg.Rule).Process(msg.Data)
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	router.POST("/expression", func(c *gin.Context) {
		var msg ExpressionMessage
		if err := c.ShouldBindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result, err := R.NewValue(msg.Expression).Evaluate(msg.Data)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		log.Printf("Starting server on :%s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}
