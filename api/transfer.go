package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/vinishs59/simplebank/db/sqlc"
)

type transferAccountRequest struct {
	FromAccount int32  `json:"from_account" binding:"required" `
	ToAccount   int32  `json:"to_account" binding:"required" `
	Amount      int32  `json:"amount" binding:"required,min=1"`
	Currency    string `json:"currency" binding:"required,oneof= USD INR"`
}

func (s *Server) transferAccount(ctx *gin.Context) {
	var req transferAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !s.ValidateAccount(ctx, int64(req.FromAccount), req.Currency) {

		return

	}

	if !s.ValidateAccount(ctx, int64(req.ToAccount), req.Currency) {

		return

	}

	arg := db.TransferTxParams{
		FromAccountID: int64(req.FromAccount),
		ToAccountID:   int64(req.ToAccount),
		Amount:        int64(req.Amount),
	}
	log.Print("Create Account")
	result, err := s.store.TransferTx(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (s *Server) ValidateAccount(ctx *gin.Context, accountID int64, currency string) bool {

	account, err := s.store.GetAccount(ctx, int32(accountID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if account.Currency.String != currency {

		err := fmt.Errorf("account %d currency mismatch %s -> %s ", account.UserID, account.Currency.String, currency)
		ctx.JSON(http.StatusBadRequest, err)
		return false

	}

	return true

}
