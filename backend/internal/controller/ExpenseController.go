package controller

import (
	customError "backend/internal/error"
	"backend/internal/handler"
	"backend/internal/model"
	"backend/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterExpenseRoutes(router *gin.Engine, repo repository.ExpenseRepo) {
	router.GET("api/v1/expense/:id", func(ctx *gin.Context) { GetExpense(ctx, repo) })
	router.GET("api/v1/expense", func(ctx *gin.Context) { ListExpenses(ctx, repo) })
	router.POST("api/v1/expense", func(ctx *gin.Context) { CreateExpense(ctx, repo) })
	router.PUT("api/v1/expense/:id", func(ctx *gin.Context) { UpdateExpense(ctx, repo) })
	router.DELETE("api/v1/expense/:id", func(ctx *gin.Context) { DeleteExpense(ctx, repo) })
}

func GetExpense(c *gin.Context, repo repository.ExpenseRepo) {
	result, err := handler.GetExpenseHandler(c, repo)
	if err != nil {
		switch err.(type) {
		case *customError.ValidationError:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

func ListExpenses(c *gin.Context, repo repository.ExpenseRepo) {
	result, err := handler.ListExpenseHandler(c, repo)
	if err != nil {
		switch err.(type) {
		case *customError.ValidationError:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func CreateExpense(c *gin.Context, repo repository.ExpenseRepo) {

	var expense model.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Set("expense", expense)
	result, err := handler.CreateExpenseHandler(c, repo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": result,
	})
}

func UpdateExpense(c *gin.Context, repo repository.ExpenseRepo) {
	result, err := handler.UpdateExpenseHandler(c, repo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func DeleteExpense(c *gin.Context, repo repository.ExpenseRepo) {
	err := handler.DeleteExpenseHandler(c, repo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
