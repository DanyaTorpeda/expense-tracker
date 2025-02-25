package handler

import (
	expensetracker "expense-tracker"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createExpense(c *gin.Context) {
	userId, err := getUser(c)
	if err != nil {
		return
	}

	var input expensetracker.Expense
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Expense.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) getAllExpenses(c *gin.Context) {

}
