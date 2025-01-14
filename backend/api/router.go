package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rubenhoenle/ci-cd-lecture-project-template/model"
)

type route struct {
	method  string
	path    string
	handler gin.HandlerFunc
}

type routeFactory interface {
	getAllRoutes(Persistence) []route
}

func addFactoryRoutes(router *gin.Engine, persistence Persistence, factory routeFactory) {
	for _, r := range factory.getAllRoutes(persistence) {
		router.Handle(r.method, r.path, r.handler)
	}
}

func NewRouter(persistence Persistence) *gin.Engine {
	router := gin.Default()
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/api/load-demo", func(ctx *gin.Context) {
		persistence.CreateUser(model.User{Uuid: uuid.New(), DisplayName: "Max", Balance: 1.50})
		persistence.CreateUser(model.User{Uuid: uuid.New(), DisplayName: "Bob", Balance: -2.50})
		persistence.CreateUser(model.User{Uuid: uuid.New(), DisplayName: "Alice", Balance: 19.0})
		ctx.Status(http.StatusOK)
	})

	addFactoryRoutes(router, persistence, userRouteFactory{})
	addFactoryRoutes(router, persistence, articleRouteFactory{})
	addFactoryRoutes(router, persistence, transactionRouteFactory{})

	return router
}
