package api

import (
	db "github.com/Tboules/back_end_master/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	router := gin.Default()
	server := &Server{
		store: store,
	}

	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccountByID)
	router.GET("/accounts", server.getAccounts)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
