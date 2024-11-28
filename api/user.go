package api

import (
	"net/http"

	db "example.com/banking/db/sqlc"
	util "example.com/banking/utils"
	"github.com/gin-gonic/gin"
)

type RegisterAccountRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,ValidEmail"`
}

func (server *Server) RegisterUser(ctx *gin.Context) {
	var req RegisterAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal Server Error with request data..."})
	}

	hashedPsw, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal Server Error Error with hashing password"})
	}

	args := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPsw,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal Server Error Error with creating user"})
	}
	ctx.JSON(http.StatusOK, user)

}

func (server *Server) LoginUser(ctx *gin.Context) {

}
