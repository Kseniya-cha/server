package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	refreshStream "github.com/Kseniya-cha/server/internal/refreshStream"
	"github.com/Kseniya-cha/server/pkg/logger"
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

func (s *refreshStreamUseCase) GetId(ctx context.Context, id interface{}) (refreshStream.RefreshStreamWithNull, error) {
	switch id.(type) {
	case string:
		val, err := strconv.Atoi(id.(string))
		if err != nil {
			return refreshStream.RefreshStreamWithNull{}, fmt.Errorf(refreshStream.ConvertIDErrConst)
		} else {
			logger.LogDebug(s.log, refreshStream.ConvertIdIntConst)
			return s.Repo.GetId(ctx, val)
		}
	default:
		return refreshStream.RefreshStreamWithNull{}, fmt.Errorf(refreshStream.ConvertIDErrConst)
	}
}

func (s *refreshStreamUseCase) Delete(ctx context.Context, id interface{}) error {
	switch id.(type) {
	case string:
		val, err := strconv.Atoi(id.(string))
		if err != nil {
			return fmt.Errorf(refreshStream.ConvertIDErrConst)
		} else {
			logger.LogDebug(s.log, refreshStream.ConvertIdIntConst)
			return s.Repo.Delete(ctx, val)
		}
	default:
		return fmt.Errorf(refreshStream.ConvertIDErrConst)
	}
}

func (s *refreshStreamUseCase) Insert(ctx context.Context, rs *refreshStream.RefreshStreamWithNull) error {
	return s.Repo.Insert(ctx, rs)
}

func (s *refreshStreamUseCase) Update(ctx context.Context, rs *refreshStream.RefreshStreamWithNull) error {
	return s.Repo.Update(ctx, rs)
}
