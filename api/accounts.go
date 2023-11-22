package api

import (
	"context"
	"net/http"

	db "github.com/Tboules/back_end_master/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type createAccountParams struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (s *Server) createAccount(c *gin.Context) {
	var req createAccountParams
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	newAccount, err := s.store.CreateAccount(context.Background(), db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  0,
		Currency: req.Currency,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, newAccount)
}

type getAccountByIDParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) getAccountByID(c *gin.Context) {
	var req getAccountByIDParams

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := s.store.GetAccount(context.Background(), req.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, account)
}

type getAccountsQueryParams struct {
	PageID   int64 `form:"page_id" binding:"required"`
	PageSize int64 `form:"page_size" binding:"required,max=50"`
}

func (s *Server) getAccounts(c *gin.Context) {
	var query getAccountsQueryParams

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	accounts, err := s.store.ListAccounts(context.Background(), db.ListAccountsParams{
		Limit:  query.PageSize,
		Offset: (query.PageID - 1) * query.PageSize,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//check if accounts are empty at that pageid
	if len(accounts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Accounts Found"})
		return
	}

	c.JSON(http.StatusOK, accounts)
}
