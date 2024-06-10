package repository

import (
	"backend/internal/model"
	"backend/internal/utils"
	"sync"

	"gorm.io/gorm"
)

// 定義一個 UserRepo 的 interface
type UserRepo interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

var (
	userRepoOnce     sync.Once
	userRepoInstance UserRepo
)

// CreateUser 用來新增一筆 User
func (r *userRepo) CreateUser(user *model.User) (*model.User, error) {
	password := user.Password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByUsernameAndPassword 用來取得指定 Username & Password 的 User
func (r *userRepo) GetUserByUsername(username string) (*model.User, error) {
	user := model.User{}
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUserRepo 用來建立一個 UserRepo 的實例
func CreateUserRepo(db *gorm.DB) UserRepo {
	userRepoOnce.Do(func() {
		userRepoInstance = &userRepo{
			db: db,
		}
	})

	return userRepoInstance
}
