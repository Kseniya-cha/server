package usecase

import (
	"context"
	"database/sql"

	refreshStream "github.com/Kseniya-cha/server/internal/refreshStream"
	"github.com/sirupsen/logrus"
)

type refreshStreamUseCase struct {
	db   *sql.DB
	log  *logrus.Logger
	Repo refreshStream.RefreshStreamRepository // интерфейс

}

func NewRefreshStreamUseCase(Repo refreshStream.RefreshStreamRepository, db *sql.DB, log *logrus.Logger) *refreshStreamUseCase {
	return &refreshStreamUseCase{
		db:   db,
		log:  log,
		Repo: Repo,
	}
}

func (s *refreshStreamUseCase) Get(ctx context.Context) ([]refreshStream.RefreshStreamWithNull, error) {
	return s.Repo.Get(ctx)
}

func (s *refreshStreamUseCase) GetId(ctx context.Context, value interface{}) (refreshStream.RefreshStreamWithNull, error) {
	return s.Repo.GetId(ctx, value)
}

func (s *refreshStreamUseCase) Delete(ctx context.Context, value int) error {
	return s.Repo.Delete(ctx, value)
}

func (s *refreshStreamUseCase) Insert(ctx context.Context, cols, vals string) error {
	return s.Repo.Insert(ctx, cols, vals)
}

func (s *refreshStreamUseCase) Update(ctx context.Context, col string, val, id interface{}) error {
	return s.Repo.Update(ctx, col, val, id)
}
