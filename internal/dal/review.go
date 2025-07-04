package dal

import (
	"context"
	"github.com/arnab-xyz/MovieVerse.git/internal/model"
	"gorm.io/gorm"
)

type ReviewDal struct {
	db *gorm.DB
}

func NewReviewDal(db *gorm.DB) *ReviewDal {
	return &ReviewDal{
		db: db,
	}
}

func (m *ReviewDal) Add(ctx context.Context, review *model.Review) error {
	txn := m.db.Create(review)
	return txn.Error
}

func (m *ReviewDal) Update(ctx context.Context, updatedReview *model.Review) error {
	txn := m.db.Save(updatedReview)
	return txn.Error
}

func (m *ReviewDal) Delete(ctx context.Context, id string) error {
	txn := m.db.Delete(&model.Review{}, "id = ?", id)
	return txn.Error
}
