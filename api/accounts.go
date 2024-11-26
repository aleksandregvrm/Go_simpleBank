package api

import (
	"database/sql"
	"net/http"

	db "example.com/banking/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (server *Server) CreateAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleDatabaseError(ctx, err)
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	accounts, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		handleDatabaseError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

type ListAccountParams struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListAccounts(ctx *gin.Context) {
	var req ListAccountParams
	if err := ctx.ShouldBindQuery(&req); err != nil {
		handleDatabaseError(ctx, err)
		return
	}

	arg := db.ListAllAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	account, err := server.store.ListAllAccounts(ctx, arg)
	if err != nil {
		handleDatabaseError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID    int64 `uri:"id" binding:"required",min=1`
	Owner string
}

func (server *Server) GetAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		handleDatabaseError(ctx, err)
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			handleDatabaseError(ctx, err)
			return
		}
		handleDatabaseError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, account)
}

type updateAccountRequest struct {
	ID      int64 `json:"id" binding:"required"`
	Balance int64 `json:"balance" binding:"required"`
}

func (server *Server) UpdateAccount(ctx *gin.Context) {
	var req updateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleDatabaseError(ctx, err)
		return
	}
	args := db.UpdateAccountParams{
		ID:      req.ID,
		Balance: req.Balance,
	}
	account, err := server.store.UpdateAccount(ctx, args)
	if err = ctx.ShouldBindJSON(&req); err != nil {
		handleDatabaseError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"acc": account, "msg": "This account has been updated"})
}
