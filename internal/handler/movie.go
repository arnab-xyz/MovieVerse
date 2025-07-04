package handler

import (
	"github.com/arnab-xyz/MovieVerse.git/internal/model"
	"github.com/arnab-xyz/MovieVerse.git/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MovieHandler struct {
	movieService *service.MovieService
}

func NewMovieHandler(movieService *service.MovieService) *MovieHandler {
	return &MovieHandler{
		movieService: movieService,
	}
}

func (m *MovieHandler) Add(ctx *gin.Context) {
	movie := &model.Movie{}
	err := ctx.BindJSON(movie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := m.movieService.Add(ctx, movie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}
	ctx.JSON(http.StatusCreated, resp)
}

func (m *MovieHandler) GetAll(ctx *gin.Context) {
	movies, err := m.movieService.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, movies)
}

func (m *MovieHandler) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	movie, err := m.movieService.Get(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, movie)
}
