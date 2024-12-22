package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "example.com/banking/db/sqlc"
	util "example.com/banking/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createAccountRequest struct {
	Currency string `json:"currency" binding:"required,ValidCurrency"`
}

func (server *Server) CreateAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleDatabaseError(ctx, err)
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*util.Payload)
	arg := db.CreateAccountParams{
		Owner:    authPayload.Username,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, gin.H{"msg": "You cannot create an account with non existing user"})
			}
		}
		handleDatabaseError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, account)
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
	authPayload := ctx.MustGet(authorizationPayloadKey).(*util.Payload)
	arg := db.ListAllAccountsParams{
		Owner:  authPayload.Username,
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
	ID int64 `uri:"id" binding:"required",min=1`
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
	authPayload := ctx.MustGet(authorizationPayloadKey).(*util.Payload)
	if account.Owner != authPayload.Username {
		err := errors.New("account doesn't belong to the authenticated user")
		handleDatabaseError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, account)
}

type updateAccountRequest struct {
	ID      int64 `json:"id" binding:"required"`
	Balance int64 `json:"balance" binding:"required"`
}

func (server *Server) updateAccount(ctx *gin.Context) {
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
	if err != nil {
		handleDatabaseError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"acc": account, "msg": "This account has been updated"})
}
