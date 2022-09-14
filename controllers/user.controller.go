package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/wpcodevo/golang-postgresql-api/db/sqlc"
	"github.com/wpcodevo/golang-postgresql-api/models"
)

type UserController struct {
	db  *db.Queries
	ctx context.Context
}

func NewUserController(db *db.Queries, ctx context.Context) UserController {
	return UserController{db, ctx}
}

func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(db.User)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": models.FilteredResponse(currentUser)}})
}
