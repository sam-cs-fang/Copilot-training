package handler

import (
	"backend/internal/model"
	"backend/internal/repository"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetExpenseHandler(c *gin.Context, repo repository.ExpenseRepo) (model.ExpenseDto, error) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return model.ExpenseDto{}, fmt.Errorf("invalid user id: %v", err)
	}

	expenseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return model.ExpenseDto{}, fmt.Errorf("invalid expense id: %v", err)
	}

	expense, err := repo.GetExpense(userId, expenseId)
	if err != nil {
		return model.ExpenseDto{}, fmt.Errorf("failed to get expense: %v", err)
	}

	return model.ExpenseDto{
		ID:       int(expense.ID),
		Title:    expense.Title,
		Amount:   expense.Amount,
		Category: expense.Category,
		Date:     expense.CreatedAt,
	}, nil
}
