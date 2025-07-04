package service

import (
	"context"
	"github.com/arnab-xyz/MovieVerse.git/internal/dal"
	"github.com/arnab-xyz/MovieVerse.git/internal/model"
	"github.com/google/uuid"
)

type MovieService struct {
	movieDal *dal.MovieDal
}

func NewMovieService(movieDal *dal.MovieDal) *MovieService {
	return &MovieService{
		movieDal: movieDal,
	}
}

func (m *MovieService) Add(ctx context.Context, movie *model.Movie) (*model.Movie, error) {
	id := uuid.New().String()
	movie.ID = id
	err := m.movieDal.Add(ctx, movie)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (m *MovieService) Get(ctx context.Context, id string) (*model.Movie, error) {
	movie, err := m.movieDal.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (m *MovieService) GetAll(ctx context.Context) ([]*model.Movie, error) {
	movies, err := m.movieDal.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
