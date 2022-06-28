package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/wpcodevo/golang-postgresql-grpc/db/sqlc"
	"github.com/wpcodevo/golang-postgresql-grpc/utils"
)

type AuthController struct {
	db *db.Queries
}

func NewAuthController(db *db.Queries) *AuthController {
	return &AuthController{db}
}

func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var credentials *db.User

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword := utils.HashPassword(credentials.Password)

	args := &db.CreateUserParams{
		Name:      credentials.Name,
		Email:     credentials.Email,
		Password:  hashedPassword,
		Photo:     "default.jpeg",
		Verified:  true,
		Role:      "user",
		UpdatedAt: time.Now(),
	}

	user, err := ac.db.CreateUser(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": user}})
}
