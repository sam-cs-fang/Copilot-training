package repository

// 此檔案用來實作與資料庫的互動邏輯

import (
	"backend/internal/model"

	"sync"

	"gorm.io/gorm"
)

// 定義一個 ExpenseRepo 的 interface
type ExpenseRepo interface {
	GetExpense(userId int, expenseId int) (*model.Expense, error)
	ListExpenses(userId int) ([]model.Expense, error)
	CreateExpense(expense *model.Expense) (*model.Expense, error)
	UpdateExpense(expenseId int, userId int, expense *model.Expense) (*model.Expense, error)
	DeleteExpense(expenseId int, userId int) error
}

// 定義一個 ExpenseRepo 的 struct
type expenseRepo struct {
	db *gorm.DB
}

var (
	expenseRepoInstance ExpenseRepo
	once                sync.Once
)

// GetExpense 用來取得指定 UserId & ExpenseId 的 Expense
func (r *expenseRepo) GetExpense(userId int, expenseId int) (*model.Expense, error) {
	expense := model.Expense{}
	err := r.db.Where("id = ? AND user_id = ?", expenseId, userId).First(&expense).Error
	if err != nil {
		return nil, err
	}

	return &expense, nil
}

// ListExpenses 用來取得所有的 Expense
func (r *expenseRepo) ListExpenses(userId int) ([]model.Expense, error) {
	expenses := []model.Expense{}
	err := r.db.Where("user_id = ?", userId).Find(&expenses).Error
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

// CreateExpense 用來新增一筆 Expense
func (r *expenseRepo) CreateExpense(expense *model.Expense) (*model.Expense, error) {
	if err := r.db.Create(expense).Error; err != nil {
		return nil, err
	}
	return expense, nil
}

// UpdateExpense 用來更新一筆 Expense
func (r *expenseRepo) UpdateExpense(expenseId int, userId int, expense *model.Expense) (*model.Expense, error) {
	var existingExpense model.Expense
	if err := r.db.Where("id = ? AND user_id = ?", expenseId, userId).First(&existingExpense).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&existingExpense).Updates(expense).Error; err != nil {
		return nil, err
	}

	return &existingExpense, nil
}

// DeleteExpense 用來刪除指定 id 的 Expense
func (r *expenseRepo) DeleteExpense(id int, userId int) error {
	if err := r.db.Where("id = ? AND user_id = ?", id, userId).Delete(&model.Expense{}).Error; err != nil {
		return err
	}

	return nil
}

// CreateExpenseRepo 用來建立一個 ExpenseRepo 的實例
func CreateExpenseRepo(db *gorm.DB) ExpenseRepo {
	once.Do(func() {
		expenseRepoInstance = &expenseRepo{
			db: db,
		}
	})

	return expenseRepoInstance
}
