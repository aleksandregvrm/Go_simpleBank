package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	db "example.com/banking/db/sqlc"
	utils "example.com/banking/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type RegisterAccountRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,ValidPassword"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,ValidEmail"`
}

type UserResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	CreatedAt         time.Time `json:"created_at"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
}

func newUserResponse(user db.User) UserResponse {
	return UserResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

func (server *Server) RegisterUser(ctx *gin.Context) {
	var req RegisterAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleDatabaseError(ctx, err)
		return
	}

	hashedPsw, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal Server Error Error with hashing password"})
		return
	}

	args := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPsw,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, args)
	if err != nil {
		handleDatabaseError(ctx, err)
		return
	}
	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Registration successful", "user": rsp})

}

type LoginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user"`
}

func (server *Server) LoginUser(ctx *gin.Context) {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	var req LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid request data", "error": err.Error()})
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "User not found", "error": err.Error()})
		return
	}

	// Verify the password
	if !utils.IsPasswordMatch(req.Password, user.HashedPassword) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid credentials"})
		return
	}

	duration, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_DURATION"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Invalid token duration", "error": err.Error()})
		return
	}
	fmt.Println(duration)
	accessToken, err := server.tokenMaker.CreateToken(user.Username, duration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to generate access token", "error": err.Error()})
		return
	}
	fmt.Println(accessToken)
	// Prepare the response
	rsp := LoginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Login successful", "data": rsp})
}

type DeleteUserParams struct {
	Username string `json:"username" binding:"required"`
}

func (server *Server) DeleteUser(ctx *gin.Context) {
	var req DeleteUserParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid request data", "error": err.Error()})
		return
	}
	accountUsers, err := server.store.GetUserWithAccounts(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid request data", "error": err.Error()})
		return
	}
	for _, user := range accountUsers {
		if user.Balance.Valid && user.Balance.Int64 > 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid request data", "error": err.Error()})
			// this for loop will check if the
			return
		}
	}

	if err := server.store.DeleteUser(ctx, req.Username); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to delete the user", "error": err.Error()})
		return
	}
	fmt.Println("user deleted")
	ctx.JSON(http.StatusOK, gin.H{"msg": "user deleted Successfully"})
}
