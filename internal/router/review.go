package router

import (
	"github.com/arnab-xyz/MovieVerse.git/internal/handler"
	"github.com/gin-gonic/gin"
)

type ReviewRouter struct {
	reviewHandler *handler.ReviewHandler
}

func NewReviewRouter(reviewHandler *handler.ReviewHandler) *ReviewRouter {
	return &ReviewRouter{
		reviewHandler: reviewHandler,
	}
}

func (m *ReviewRouter) CreateRouter(server *gin.Engine) {
	server.POST("/reviews", m.reviewHandler.Add)
	server.PUT("/reviews/:id", m.reviewHandler.Update)
	server.DELETE("/reviews/:id", m.reviewHandler.Delete)
}
