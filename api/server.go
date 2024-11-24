package api

import (
	"errors"
	"net/http"

	db "example.com/banking/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// Serving http requests
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts/:id", server.GetAccount)
	router.GET("/accounts", server.ListAccounts)

	server.router = router
	return server
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