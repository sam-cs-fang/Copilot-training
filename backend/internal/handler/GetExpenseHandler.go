package handler

import (
	customError "backend/internal/error"
	"backend/internal/model"
	"backend/internal/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetExpenseHandler(c *gin.Context, repo repository.ExpenseRepo) (model.ExpenseDto, error) {
	userId, exists := c.Get("userId")
	if !exists {
		return model.ExpenseDto{}, &customError.ValidationError{Message: "UserId not found in context"}
	}

	expenseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return model.ExpenseDto{}, &customError.ValidationError{Message: "Invalid expense id"}
	}

	expense, err := repo.GetExpense(userId.(int), expenseId)
	if err != nil {
		return model.ExpenseDto{}, customError.HandleGormError(err)
	}

	return model.ExpenseDto{
		ID:       int(expense.ID),
		Title:    expense.Title,
		Amount:   expense.Amount,
		Category: expense.Category,
		Date:     expense.CreatedAt,
	}, nil
}
