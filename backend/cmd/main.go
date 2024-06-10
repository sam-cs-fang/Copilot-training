package main

import (
	"backend/config"
	"backend/internal/controller"
	"backend/internal/middleware"
	"backend/internal/model"
	"backend/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	config.LoadConfig()

	r := gin.Default()

	r.Use(middleware.Logger())
	r.Use(middleware.Auth())

	postgresDb := model.CreateDatabase()
	expenseRepo := repository.CreateExpenseRepo(postgresDb)

	// 註冊路由
	controller.RegisterExpenseRoutes(r, expenseRepo)

	// 啟動服務
	port := viper.GetString("server.port")
	r.Run(":" + port)
}
