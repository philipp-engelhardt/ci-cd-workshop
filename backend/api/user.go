package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rubenhoenle/ci-cd-lecture-project-template/model"
)

type userRouteFactory struct{}

func (factory userRouteFactory) getAllRoutes(persistence Persistence) []route {
	return []route{
		{
			// get all users
			method: http.MethodGet,
			path:   "/api/user",
			handler: func(ctx *gin.Context) {
				users, err := persistence.GetAllUsers()
				if err != nil {
					ctx.AbortWithError(http.StatusInternalServerError, err)
					return
				}
				ctx.IndentedJSON(http.StatusOK, users)
			},
		},
		{
			// create user
			method: http.MethodPost,
			path:   "/api/user",
			handler: func(ctx *gin.Context) {
				var newUser model.User
				if err := ctx.BindJSON(&newUser); err != nil {
					ctx.AbortWithError(http.StatusBadRequest, err)
					return
				}

				// generate and set a uuid
				newUuid, err := uuid.NewUUID()
				if err != nil {
					ctx.AbortWithError(http.StatusInternalServerError, err)
					return
				}
				newUser.Uuid = newUuid

				err = persistence.CreateUser(newUser)
				if err != nil {
					ctx.AbortWithError(http.StatusInternalServerError, err)
					return
				}
				ctx.IndentedJSON(http.StatusCreated, newUser)
			},
		},
		{
			// update user
			method: http.MethodPut,
			path:   "/api/user/:uuid",
			handler: func(ctx *gin.Context) {
				uuidStr := ctx.Param("uuid")
				parsedUuid, err := uuid.Parse(uuidStr)
				if err != nil {
					ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid uuid"})
					return
				}

				var updatedUser model.User
				if err := ctx.BindJSON(&updatedUser); err != nil {
					ctx.AbortWithError(http.StatusBadRequest, err)
					return
				}

				// check if a user with the uuid exists
				fetchedUser, err := persistence.GetUser(parsedUuid)
				if err != nil {
					ctx.AbortWithError(http.StatusInternalServerError, err)
					return
				}
				if fetchedUser == nil {
					ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found"})
					return
				}

				// prevent change of uuid
				updatedUser.Uuid = parsedUuid

				err = persistence.UpdateUser(updatedUser)
				if err != nil {
					ctx.AbortWithError(http.StatusInternalServerError, err)
					return
				}
				ctx.IndentedJSON(http.StatusOK, updatedUser)
			},
		},
		{
			// get user
			method: http.MethodGet,
			path:   "/api/user/:uuid",
			handler: func(ctx *gin.Context) {
				uuidStr := ctx.Param("uuid")
				parsedUuid, err := uuid.Parse(uuidStr)
				if err != nil {
					ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid uuid"})
					return
				}

				fetchedUser, err := persistence.GetUser(parsedUuid)
				if err != nil {
					ctx.AbortWithError(http.StatusInternalServerError, err)
					return
				}
				// check if a user with the uuid exists
				if fetchedUser == nil {
					ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found"})
					return
				}

				ctx.IndentedJSON(http.StatusOK, fetchedUser)
			},
		},
		{
			// delete user
			method: http.MethodDelete,
			path:   "/api/user/:uuid",
			handler: func(ctx *gin.Context) {
				uuidStr := ctx.Param("uuid")
				parsedUuid, err := uuid.Parse(uuidStr)
				if err != nil {
					ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid uuid"})
					return
				}

				// check if a user with the uuid exists
				fetchedUser, err := persistence.GetUser(parsedUuid)
				if err != nil {
					ctx.AbortWithError(http.StatusInternalServerError, err)
					return
				}
				if fetchedUser == nil {
					ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found"})
					return
				}

				err = persistence.DeleteUser(parsedUuid)
				if err != nil {
					ctx.AbortWithError(http.StatusInternalServerError, err)
					return
				}
				ctx.IndentedJSON(http.StatusOK, gin.H{"message": "OK"})
			},
		},
	}
}
