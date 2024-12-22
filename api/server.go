package api

import (
	"errors"
	"net/http"
	"os"

	db "example.com/banking/db/sqlc"
	utils "example.com/banking/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

// Serving http requests
type Server struct {
	store      *db.Store
	tokenMaker utils.Maker
	router     *gin.Engine
}

// Creates server with provided routes and env vars applied
func NewServer(store *db.Store) (*Server, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	tokenMaker, err := utils.NewPasetoMaker(os.Getenv("TOKEN_SYMMETRIC_KEY"))
	server := &Server{store: store, tokenMaker: tokenMaker}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("ValidCurrency", ValidCurrency)
		v.RegisterValidation("ValidEmail", ValidEmail)
		v.RegisterValidation("ValidPassword", ValidPassword)
	}

	server.registerRouters(router)

	server.router = router
	return server, nil
}

func (server *Server) registerRouters(router *gin.Engine) {

	// creating the authorized only route
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.CreateAccount)
	authRoutes.GET("/accounts/:id", server.GetAccount)
	authRoutes.GET("/accounts", server.ListAccounts)
	authRoutes.POST("/transfers", server.CreateTransfer)
	router.POST("/user/register", server.RegisterUser)
	router.POST("/user/login", server.LoginUser)
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func handleDatabaseError(ctx *gin.Context, err error) {
	if pgErr, ok := err.(*pq.Error); ok {
		switch pgErr.Code.Name() {
		case "unique_violation":
			ctx.JSON(http.StatusConflict, errorResponse(errors.New("account already exists")))
		case "foreign_key_violation":
			ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("invalid foreign key")))
		default:
			ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("database error")))
		}
		return
	}

	// Fallback for unknown errors
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
}
