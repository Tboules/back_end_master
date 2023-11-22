package api

import (
	"context"
	"net/http"

	db "github.com/Tboules/back_end_master/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountParams struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (s *Server) createAccount(ctx *gin.Context) {
	var req createAccountParams
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorRresponse(err))
	}

	newAccount, err := s.store.CreateAccount(context.Background(), db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  0,
		Currency: req.Currency,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorRresponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newAccount)
}
