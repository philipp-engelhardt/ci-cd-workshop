package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionRouteFactory struct{}

func (factory transactionRouteFactory) getAllRoutes(persistence Persistence) []route {
	return []route{
		{
			method: http.MethodGet,
			path:   "/api/transaction",
			handler: func(ctx *gin.Context) {
				transactions, err := persistence.GetAllTransactions()
				if err != nil {
					ctx.AbortWithError(http.StatusInternalServerError, err)
					return
				}
				ctx.IndentedJSON(http.StatusOK, transactions)
			},
		},
	}
}
