package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-postgresql-api/controllers"
	db "github.com/wpcodevo/golang-postgresql-api/db/sqlc"
	"github.com/wpcodevo/golang-postgresql-api/middleware"
)

type UserRoutes struct {
	userController controllers.UserController
	db             *db.Queries
}

func NewUserRoutes(userController controllers.UserController, db *db.Queries) UserRoutes {
	return UserRoutes{userController, db}
}

func (rc *UserRoutes) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("/users")
	router.GET("/me", middleware.DeserializeUser(rc.db), rc.userController.GetMe)
}
