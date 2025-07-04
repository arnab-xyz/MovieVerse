package dal

import (
	"context"
	"github.com/arnab-xyz/MovieVerse.git/internal/model"
	"gorm.io/gorm"
)

type MovieDal struct {
	db *gorm.DB
}

func NewMovieDal(db *gorm.DB) *MovieDal {
	return &MovieDal{
		db: db,
	}
}

func (m *MovieDal) Add(ctx context.Context, movie *model.Movie) error {
	txn := m.db.Create(movie)
	return txn.Error
}

func (m *MovieDal) Get(ctx context.Context, id string) (*model.Movie, error) {
	movie := &model.Movie{}
	txn := m.db.Preload("Reviews").
		First(&movie, "id = ?", id)
	if txn.Error != nil {
		return nil, txn.Error
	}
	return movie, nil
}

func (m *MovieDal) GetAll(ctx context.Context) ([]*model.Movie, error) {
	movie := []*model.Movie{}
	txn := m.db.Find(&movie)
	if txn.Error != nil {
		return nil, txn.Error
	}
	return movie, nil
}
