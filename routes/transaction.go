package routes

import (
	"github.com/Favoree-Team/server-user-api/controller"
	"github.com/Favoree-Team/server-user-api/repository"
	"github.com/Favoree-Team/server-user-api/service"

	"github.com/gin-gonic/gin"
)

var (
	transRepo       = repository.NewTransactionRepository(DB)
	transService    = service.NewTransactionRepository(transRepo)
	transController = controller.NewTransactionController(transService)
)

func TransactionRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		trx := v1.Group("/transactions")
		{

			// ================  ADMIN  =====================
			// get transaction list, admin only
			trx.GET("/:page/:limit", mainMiddleware, transController.GetAllTransaction)

			// update transaction status flow, admin only
			trx.PUT("/:transaction_id", mainMiddleware, transController.UpdateTransactionByAdmin)
			/*
				body {
					"status": "paid"
				}
			*/

			// ================  USER  =====================
			// create new transaction
			trx.POST("/", mainMiddleware, transController.CreateTransaction)

			// get transaction detail (MASIH BELUM PERLU)
			// bisa diget dari last transaction
			//trx.GET("/:transaction_id", mainMiddleware)

			// last transaction user
			trx.GET("/last", mainMiddleware, transController.GetLastTransaction)

			// user post after click confirm paid
			trx.POST("/:transaction_id/confirm_paid", mainMiddleware, transController.ConfirmPaid)

			// cancel by user, with middleware
			trx.POST("/:transaction_id/cancel", mainMiddleware, transController.CancelTransaction)
		}
	}
}
