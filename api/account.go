package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/vinishs59/simplebank/db/sqlc"
)

type createAccountRequest struct {
	OwnerName string         `json:"owner_name" binding:"required" `
	Currency  string         `json:"currency" binding:"required,oneof= USD INR"`
}

func (s *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req);err!= nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var currency sql.NullString
	if req.Currency != "" {
    currency = sql.NullString{String: req.Currency, Valid: true}
} else {
    currency = sql.NullString{Valid: false}
}

	arg := db.CreateAccountParams{
		OwnerName: req.OwnerName,
		Currency: currency,
		Balance: 0,
	}
log.Print("Create Account")
	account,err :=s.store.CreateAccount(ctx,arg)

	if err != nil {
		 ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		 return
	}

	ctx.JSON(http.StatusOK,account)
}