package handler

import (
	"backend/internal/controller/requests"
	"backend/internal/model"
	"backend/internal/repository"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Implement the business logic of CreateUser
// receive an User
// Create a new user in the database
// return UserDto
func UserSignUpHandler(c *gin.Context, repo repository.UserRepo) (model.UserDto, error) {
	user, exist := c.Get("user")
	if !exist {
		return model.UserDto{}, fmt.Errorf("failed to get user data from context")
	}

	userRequest := user.(*requests.UserSignInAndSignUpRequest)
	userData := model.User{
		Username: userRequest.Username,
		Password: userRequest.Password,
	}

	result, err := repo.CreateUser(&userData)
	if err != nil {
		return model.UserDto{}, fmt.Errorf("failed to create user: %v", err)
	}

	return model.UserDto{
		ID: result.ID,
	}, nil
}
