package handler

import (
	"backend/internal/model"
	"backend/internal/repository"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Implement the business logic of UpdateExpense
// receive an Expense
// Update an expense in the database
// return nothing
func UpdateExpenseHandler(c *gin.Context, repo repository.ExpenseRepo) (model.ExpenseDto, error) {
	expense, exist := c.Get("expense")
	if !exist {
		return model.ExpenseDto{}, fmt.Errorf("failed to get expense from context")
	}

	result, err := repo.UpdateExpense(expense.(*model.Expense))
	if err != nil {
		return model.ExpenseDto{}, fmt.Errorf("failed to update expense: %v", err)
	}

	return model.ExpenseDto{
		ID:       int(result.ID),
		Title:    result.Title,
		Amount:   result.Amount,
		Category: result.Category,
		Date:     result.CreatedAt,
	}, nil

}
