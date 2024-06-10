package controller

import (
	"backend/internal/controller/requests"
	customError "backend/internal/error"
	handler "backend/internal/handler/user"
	"backend/internal/repository"
	"backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, repo repository.UserRepo) {
	router.POST("api/v1/users/signin", func(ctx *gin.Context) { UserSignIn(ctx, repo) })
	router.POST("api/v1/users/signup", func(ctx *gin.Context) { UserSignUp(ctx, repo) })
}

// 使用者登入，登入成功後會將 userId 包裹在 JWT 中回傳
func UserSignIn(c *gin.Context, repo repository.UserRepo) {
	var user *requests.UserSignInAndSignUpRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Set("user", user)
	result, err := handler.UserSignInHandler(c, repo)
	if err != nil {
		switch err.(type) {
		case *customError.ValidationError:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		case *customError.NotFoundError:
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	// 產生 JWT Token，並將 userId 包裹在 JWT 中回傳
	token, err := utils.GenerateJWT(result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": token,
	})
}

// 使用者註冊，註冊成功後會將 userId 包裹在 JWT 中回傳
func UserSignUp(c *gin.Context, repo repository.UserRepo) {
	var user *requests.UserSignInAndSignUpRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Set("user", user)
	result, err := handler.UserSignUpHandler(c, repo)
	if err != nil {
		switch err.(type) {
		case *customError.ValidationError:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		case *customError.NotFoundError:
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	// 產生 JWT Token，並將 userId 包裹在 JWT 中回傳
	token, err := utils.GenerateJWT(result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": token,
	})
}
