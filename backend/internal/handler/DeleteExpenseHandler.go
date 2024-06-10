package handler

import (
	customError "backend/internal/error"
	"backend/internal/repository"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Implement the business logic of DeleteExpense
// receive an Expense
// Delete an expense in the database
// return nothing
func DeleteExpenseHandler(c *gin.Context, repo repository.ExpenseRepo) error {
	expenseId := c.Param("id")
	expenseIdInt, err := strconv.Atoi(expenseId)
	if err != nil {
		return &customError.ValidationError{Message: fmt.Sprintf("invalid expense id: %v", err)}
	}

	userId, exists := c.Get("userId")
	if !exists {
		return &customError.ValidationError{Message: "UserId not found in context"}
	}

	err = repo.DeleteExpense(expenseIdInt, int(userId.(float64)))
	if err != nil {
		return fmt.Errorf("failed to delete expense: %v", err)
	}

	return nil
}
