package handler

import (
	"errors"
	expensetracker "expense-tracker"
	"net/http"
	"strconv"

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
	id, err := getUser(c)
	if err != nil {
		return
	}

	sorting := c.Query("sort_by")
	order := c.Query("order")

	expenses, err := h.service.Expense.GetAll(id, sorting, order)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, expenses)
}

func (h *Handler) getById(c *gin.Context) {
	userId, err := getUser(c)
	if err != nil {
		return
	}

	val := c.Param("id")
	if val == "" {
		newErrorResponse(c, http.StatusBadRequest, errors.New("").Error())
		return
	}

	expenseId, err := strconv.Atoi(val)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	expense, err := h.service.Expense.GetById(userId, expenseId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, expense)
}

func (h *Handler) updateExpense(c *gin.Context) {
	userId, err := getUser(c)
	if err != nil {
		return
	}

	val := c.Param("id")
	if val == "" {
		newErrorResponse(c, http.StatusBadRequest, errors.New("").Error())
		return
	}

	expenseId, err := strconv.Atoi(val)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input expensetracker.Expense
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.Expense.Update(userId, expenseId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) deleteExpense(c *gin.Context) {
	userId, err := getUser(c)
	if err != nil {
		return
	}

	val := c.Param("id")
	if val == "" {
		newErrorResponse(c, http.StatusBadRequest, errors.New("").Error())
		return
	}

	expenseId, err := strconv.Atoi(val)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Expense.Delete(userId, expenseId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
