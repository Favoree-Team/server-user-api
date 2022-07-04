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
			trx.GET("/:page/:limit")

			// update transaction status flow, admin only
			trx.PUT("/:id")
			/*
				body {
					"status": "paid"
				}
			*/

			// ================  USER  =====================
			// create new transaction
			trx.POST("/")

			// get transaction detail
			trx.GET("/:id")

			// user post after click confirm paid
			trx.POST("/:id/confirm_paid")

			// cancel by user, with middleware
			trx.POST("/:id/cancel")
		}
	}
}
