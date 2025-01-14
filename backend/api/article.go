package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type articleRouteFactory struct{}

func (factory articleRouteFactory) getAllRoutes(persistence Persistence) []route {
	return []route{
		{
			method: http.MethodGet,
			path:   "/api/article",
			handler: func(ctx *gin.Context) {
				articles, err := persistence.GetAllArticles()
				if err != nil {
					ctx.AbortWithError(http.StatusInternalServerError, err)
					return
				}
				ctx.IndentedJSON(http.StatusOK, articles)
			},
		},
	}
}
