package main

import (
	"github.com/arnab-xyz/MovieVerse.git/internal/dal"
	"github.com/arnab-xyz/MovieVerse.git/internal/handler"
	"github.com/arnab-xyz/MovieVerse.git/internal/model"
	"github.com/arnab-xyz/MovieVerse.git/internal/router"
	"github.com/arnab-xyz/MovieVerse.git/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() {
	server := gin.Default()
	db, err := gorm.Open(sqlite.Open("MovieVerse.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&model.Movie{})
	db.AutoMigrate(&model.Review{})
	movieDal := dal.NewMovieDal(db)
	movieServices := service.NewMovieService(movieDal)
	movieHandler := handler.NewMovieHandler(movieServices)
	movieRouter := router.NewMovieRouter(movieHandler)
	movieRouter.CreateRouter(server)
	reviewDal := dal.NewReviewDal(db)
	reviewServices := service.NewReviewService(reviewDal)
	reviewHandler := handler.NewReviewHandler(reviewServices)
	reviewRouter := router.NewReviewRouter(reviewHandler)
	reviewRouter.CreateRouter(server)
	server.Run("localhost:8080")
}
