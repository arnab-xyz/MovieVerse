package handler

import (
	"github.com/arnab-xyz/MovieVerse.git/internal/model"
	"github.com/arnab-xyz/MovieVerse.git/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReviewHandler struct {
	reviewService *service.ReviewService
}

func NewReviewHandler(reviewService *service.ReviewService) *ReviewHandler {
	return &ReviewHandler{
		reviewService: reviewService,
	}
}

func (m *ReviewHandler) Add(ctx *gin.Context) {
	review := &model.Review{}
	err := ctx.BindJSON(review)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := m.reviewService.Add(ctx, review)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}
	ctx.JSON(http.StatusCreated, resp)
}

func (m *ReviewHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	updatedReview := &model.Review{}
	err := ctx.BindJSON(updatedReview)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = m.reviewService.Update(ctx, id, updatedReview)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (m *ReviewHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := m.reviewService.Delete(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
