package handler

import (
	"backend/internal/controller/requests"
	customError "backend/internal/error"
	"backend/internal/model"
	"backend/internal/repository"
	"backend/internal/utils"

	"github.com/gin-gonic/gin"
)

func UserSignInHandler(c *gin.Context, repo repository.UserRepo) (model.UserDto, error) {
	user, exist := c.Get("user")
	if !exist {
		return model.UserDto{}, &customError.ValidationError{Message: "failed to get user data from context"}
	}

	result, err := repo.GetUserByUsername(
		user.(*requests.UserSignInAndSignUpRequest).Username,
	)
	if err != nil {
		return model.UserDto{}, customError.HandleGormError(err)
	}

	password := user.(*requests.UserSignInAndSignUpRequest).Password
	isPass := utils.CheckPasswordHash(password, result.Password)
	if !isPass {
		return model.UserDto{}, &customError.UnAuthorizeError{Message: "password is incorrect"}
	}

	return model.UserDto{
		ID: result.ID,
	}, nil
}
