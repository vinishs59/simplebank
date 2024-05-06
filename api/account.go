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

type getAccountReq struct{
	ID  int32 `uri:"id" binding:"required,min=1"` 
}

func (s *Server) getAccount(ctx *gin.Context) {
	var req getAccountReq
	if err := ctx.ShouldBindUri(&req);err!= nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	account ,err := s.store.GetAccount(ctx,req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
   }
   ctx.JSON(http.StatusOK,account)

}

type listAccountReq struct{
	PageID  int32 `form:"page_id" binding:"required,min=1"` 
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"` 

}


func (s *Server) listAccount(ctx *gin.Context) {
	var req listAccountReq
	if err := ctx.ShouldBindQuery(&req);err!= nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListAccountParams{
		Limit : req.PageSize,
		Offset: (req.PageID-1) * req.PageSize,
	}
	account ,err := s.store.ListAccount(ctx,arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
   }
   ctx.JSON(http.StatusOK,account)

}

