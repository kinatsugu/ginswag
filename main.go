package main

import (
	"fmt"
	"net/http"

	_ "ginswag/docs" // 必ずdocsのパスをインポートしてください
	"ginswag/models"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Ginswag Example API
// @version 1.0
// @description This is a sample API for demonstrating Swagger with Gin.
// @host localhost:8080
// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Router /ping [get]
func PingExample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// HealthCheck godoc
// @Summary health check
// @Schemes
// @Description do health check
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// GetTodo godoc
// @Summary get todo
// @Schemes
// @Description get todo
// @Tags todo
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} models.Todo
// @Router /todo/{id} [get]
func GetTodo(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	t := models.Todo{
		ID:     1,
		Title:  "title",
		Status: "status",
	}
	c.JSON(http.StatusOK, t)
}

// main 関数
func main() {
	gin.ForceConsoleColor()
	r := gin.Default()

	// Swaggerのハンドラーを追加
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// エンドポイントの定義
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", PingExample)
		v1.GET("/health", HealthCheck)
		v1.GET("/todo/:id", GetTodo)
	}

	// サーバーの起動
	r.Run(":8080")
}
