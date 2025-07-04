package router

import (
	"github.com/arnab-xyz/MovieVerse.git/internal/handler"
	"github.com/gin-gonic/gin"
)

type MovieRouter struct {
	movieHandler *handler.MovieHandler
}

func NewMovieRouter(movieHandler *handler.MovieHandler) *MovieRouter {
	return &MovieRouter{
		movieHandler: movieHandler,
	}
}

func (m *MovieRouter) CreateRouter(server *gin.Engine) {
	server.POST("/movies", m.movieHandler.Add)
	server.GET("/movies", m.movieHandler.GetAll)
	server.GET("/movies/:id", m.movieHandler.Get)
}
