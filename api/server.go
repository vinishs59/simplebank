package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	db "github.com/vinishs59/simplebank/db/sqlc"
)


type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	router := gin.Default()
	server := &Server {
		store: store,
		router: router,
	}

	router.POST("/accounts", server.createAccount )


	return server
}

func (server *Server)Start(address string ) error{ 
	return server.router.Run(address) 

}

func errorResponse (err error) gin.H {
	return gin.H{"error":err.Error()}
}

