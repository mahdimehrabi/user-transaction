package routes

import (
	controller "bbdk/app/api/controllers"
	"bbdk/app/api/middleware"
	"bbdk/domain/service"
	"github.com/gin-gonic/gin"
)

type TransactionRouter struct {
	transactionController controller.TransactionController
	authMiddleware        middleware.AuthMiddleware
}

func NewTransactionRouter(transactionService service.TransactionService, authMiddleware middleware.AuthMiddleware) *TransactionRouter {
	transactionController := controller.NewTransactionController(transactionService)
	return &TransactionRouter{transactionController: *transactionController, authMiddleware: authMiddleware}
}

func (rh *TransactionRouter) SetupRoutes(router *gin.Engine) {
	g := router.Group("/api/transactions").Use(rh.authMiddleware.Handle())
	{
		g.GET("/report/:userID", rh.transactionController.GetTransactionReportByUserID)
		g.POST("", rh.transactionController.CreateTransaction)
		g.GET("/:id", rh.transactionController.GetTransactionByID)
		g.PUT("/:id", rh.transactionController.UpdateTransaction)
		g.DELETE("/:id", rh.transactionController.DeleteTransaction)
		g.GET("", rh.transactionController.GetAllTransactions)
	}
}
