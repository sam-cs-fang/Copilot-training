package handler

import (
	"backend/internal/model"
	"backend/internal/repository"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Implement the business logic of CreateExpense
// receive an Expense
// Create a new expense in the database
// return nothing
func CreateExpenseHandler(c *gin.Context, repo repository.ExpenseRepo) (model.ExpenseDto, error) {
	expense, exist := c.Get("expense")
	if !exist {
		return model.ExpenseDto{}, fmt.Errorf("failed to get expense from context")
	}

	result, err := repo.CreateExpense(expense.(*model.Expense))
	if err != nil {
		return model.ExpenseDto{}, fmt.Errorf("failed to create expense: %v", err)
	}

	return model.ExpenseDto{
		ID:       int(result.ID),
		Title:    result.Title,
		Amount:   result.Amount,
		Category: result.Category,
		Date:     result.CreatedAt,
	}, nil
}
