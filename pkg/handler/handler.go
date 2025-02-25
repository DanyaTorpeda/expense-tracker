package handler

import (
	"expense-tracker/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	routes := gin.New()

	auth := routes.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := routes.Group("/api", h.userIdentity)
	{
		expenses := api.Group("/expenses")
		{
			expenses.POST("/", h.createExpense)
			expenses.GET("/", h.getAllExpenses)
		}
	}

	return routes
}
