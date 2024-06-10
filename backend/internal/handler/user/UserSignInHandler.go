package handler

import (
	"backend/internal/controller/requests"
	"backend/internal/model"
	"backend/internal/repository"
	"backend/internal/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserSignInHandler(c *gin.Context, repo repository.UserRepo) (model.UserDto, error) {
	user, exist := c.Get("user")
	if !exist {
		return model.UserDto{}, fmt.Errorf("failed to get expense from context")
	}

	result, err := repo.GetUserByUsernameAndPassword(
		user.(*requests.UserSignInAndSignUpRequest).Username,
		user.(*requests.UserSignInAndSignUpRequest).Password,
	)
	if err != nil {
		return model.UserDto{}, fmt.Errorf("failed to get user: %v", err)
	}

	// 驗證密碼是否正確
	isPass := utils.CheckPasswordHash(user.(*requests.UserSignInAndSignUpRequest).Password, result.Password)
	if !isPass {
		// 密碼錯誤，回傳 401
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "password is incorrect",
		})
		return model.UserDto{}, fmt.Errorf("password is incorrect")
	}

	return model.UserDto{
		ID: result.ID,
	}, nil
}
