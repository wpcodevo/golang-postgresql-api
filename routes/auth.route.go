package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-postgresql-api/controllers"
)

type AuthRoutes struct {
	authController controllers.AuthController
}

func NewAuthRoutes(authController controllers.AuthController) AuthRoutes {
	return AuthRoutes{authController}
}

func (rc *AuthRoutes) AuthRoute(rg *gin.RouterGroup) {

	router := rg.Group("/auth")
	router.POST("/register", rc.authController.SignUpUser)
}
