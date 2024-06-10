package handler

import (
	customError "backend/internal/error"
	"backend/internal/model"
	"backend/internal/repository"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Implement the business logic of ListExpense
// receive nothing
// Get all expenses from the database as an array of Expense
// return an array of ExpenseDto

func ListExpenseHandler(c *gin.Context, repo repository.ExpenseRepo) ([]model.ExpenseDto, error) {
	userId, exist := c.Get("userId")
	if !exist {
		return []model.ExpenseDto{}, &customError.ValidationError{Message: "userId not found in context"}
	}

	expenses, err := repo.ListExpenses(int(userId.(float64)))
	if err != nil {
		return []model.ExpenseDto{}, fmt.Errorf("failed to list expenses: %v", err)
	}

	expenseDtos := make([]model.ExpenseDto, 0)
	for _, expense := range expenses {
		expenseDtos = append(expenseDtos, model.ExpenseDto{
			ID:       int(expense.ID),
			Title:    expense.Title,
			Amount:   expense.Amount,
			Category: expense.Category,
			Date:     expense.CreatedAt,
		})
	}

	return expenseDtos, nil
}
