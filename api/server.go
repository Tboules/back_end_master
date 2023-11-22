package api

import (
	db "github.com/Tboules/back_end_master/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(s *db.Store) *Server {
	r := gin.Default()

	r.POST("/account", s.createAccount)

	return &Server{
		store:  s,
		router: r,
	}
}

func errorRresponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
