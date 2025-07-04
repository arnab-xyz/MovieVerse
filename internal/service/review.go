package service

import (
	"context"
	"github.com/arnab-xyz/MovieVerse.git/internal/dal"
	"github.com/arnab-xyz/MovieVerse.git/internal/model"
	"github.com/google/uuid"
)

type ReviewService struct {
	reviewDal *dal.ReviewDal
}

func NewReviewService(reviewDal *dal.ReviewDal) *ReviewService {
	return &ReviewService{
		reviewDal: reviewDal,
	}
}

func (m *ReviewService) Add(ctx context.Context, Review *model.Review) (*model.Review, error) {
	id := uuid.New().String()
	Review.ID = id
	err := m.reviewDal.Add(ctx, Review)
	if err != nil {
		return nil, err
	}
	return Review, nil
}

func (m *ReviewService) Update(ctx context.Context, id string, updatedReview *model.Review) error {
	updatedReview.ID = id
	err := m.reviewDal.Update(ctx, updatedReview)
	return err
}

func (m *ReviewService) Delete(ctx context.Context, id string) error {
	err := m.reviewDal.Delete(ctx, id)
	return err
}
