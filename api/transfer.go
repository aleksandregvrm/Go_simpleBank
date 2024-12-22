package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	db "example.com/banking/db/sqlc"
	util "example.com/banking/utils/token"
	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,ValidCurrency"`
}

func (server *Server) CreateTransfer(ctx *gin.Context) {
	var req transferRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fromAccount, valid := server.validAccount(ctx, req.FromAccountID, req.Currency)
	if !valid {
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*util.Payload)
	if fromAccount.Owner != authPayload.Username {
		err := errors.New("You cannot send money from somebody elses account")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	_, valid = server.validAccount(ctx, req.FromAccountID, req.Currency)
	if !valid {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("Not a valid transaction")))
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	fmt.Println(result)

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) (db.Account, bool) {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
			return account, false
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch account"})
		return account, false
	}

	if account.Currency != currency {
		err := fmt.Errorf("currency mismatch: expected %s, got %s", currency, account.Currency)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return account, false
	}
	return account, true
}
