package handler

import (
	customError "backend/internal/error"
	"backend/internal/model"
	"backend/internal/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Implement the business logic of UpdateExpense
// receive an Expense
// Update an expense in the database
// return nothing
func UpdateExpenseHandler(c *gin.Context, repo repository.ExpenseRepo) (model.ExpenseDto, error) {
	expense, exist := c.Get("expense")
	if !exist {
		return model.ExpenseDto{}, &customError.ValidationError{Message: "Expense not found in context"}
	}

	userId, exists := c.Get("userId")
	if !exists {
		return model.ExpenseDto{}, &customError.ValidationError{Message: "UserId not found in context"}
	}

	expenseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return model.ExpenseDto{}, &customError.ValidationError{Message: "Invalid expense id"}
	}

	result, err := repo.UpdateExpense(expenseId, int(userId.(float64)), expense.(*model.Expense))
	if err != nil {
		return model.ExpenseDto{}, customError.HandleGormError(err)
	}

	return model.ExpenseDto{
		ID:       int(result.ID),
		Title:    result.Title,
		Amount:   result.Amount,
		Category: result.Category,
		Date:     result.CreatedAt,
	}, nil

}
