package main

import (
	"backend/config"
	"backend/internal/controller"
	"backend/internal/middleware"
	"backend/internal/model"
	"backend/internal/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	config.LoadConfig()

	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"}
	r.Use(cors.New(corsConfig))

	r.Use(middleware.Logger())

	postgresDb := model.CreateDatabase()
	expenseRepo := repository.CreateExpenseRepo(postgresDb)
	userRepo := repository.CreateUserRepo(postgresDb)

	// 要驗證 token 的 routes
	private := r.Group("/api/v1")
	private.Use(middleware.Auth())
	{
		controller.RegisterExpenseRoutes(private, expenseRepo)
	}

	// 不用驗證 token 的 routes
	public := r.Group("/api/v1")
	{
		controller.RegisterUserRoutes(public, userRepo)
	}

	// 啟動服務
	port := viper.GetString("server.port")
	r.Run(":" + port)
}
