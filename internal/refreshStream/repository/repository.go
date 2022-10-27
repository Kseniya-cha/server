package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Kseniya-cha/server/constants"
	refreshStream "github.com/Kseniya-cha/server/internal/refreshStream"
	"github.com/Kseniya-cha/server/pkg/logger"
	"github.com/sirupsen/logrus"
)

type refreshStreamRepository struct {
	db  *sql.DB
	log *logrus.Logger
}

func NewRefreshStreamRepository(db *sql.DB, log *logrus.Logger) *refreshStreamRepository {
	return &refreshStreamRepository{
		db:  db,
		log: log,
	}
}

func (s refreshStreamRepository) Get(ctx context.Context) ([]refreshStream.RefreshStreamWithNull, error) {

	logger.LogDebug(s.log, constants.GetSentRespConst)
	rows, err := s.db.QueryContext(ctx, constants.GetQueryConst)
	if err != nil {
		logger.LogError(s.log, err)
		return nil, err
	}
	defer rows.Close()

	// слайс копий структур
	refreshStreamArr := []refreshStream.RefreshStreamWithNull{}
	for rows.Next() {
		rs := refreshStream.RefreshStreamWithNull{}
		err := rows.Scan(&rs.Id, &rs.Auth, &rs.Ip, &rs.Stream,
			&rs.Run, &rs.Portsrv, &rs.Sp, &rs.Camid, &rs.Record_status,
			&rs.Stream_status, &rs.Record_state, &rs.Stream_state)
		if err != nil {
			logger.LogError(s.log, err)
			return nil, err
		}
		refreshStreamArr = append(refreshStreamArr, rs)
	}
	logger.LogDebug(s.log, constants.GetIdRespConst)
	return refreshStreamArr, nil
}

func (s refreshStreamRepository) GetId(ctx context.Context,
	value interface{}) (refreshStream.RefreshStreamWithNull, error) {

	rows, err := s.db.QueryContext(ctx, fmt.Sprintf(constants.GetIdQueryConst, "id"), value)
	if err != nil {
		logger.LogError(s.log, "")
		return refreshStream.RefreshStreamWithNull{}, err
	}
	logger.LogDebug(s.log, constants.GetSentRespConst)

	rs := refreshStream.RefreshStreamWithNull{}
	for rows.Next() {
		err = rows.Scan(&rs.Id, &rs.Auth, &rs.Ip, &rs.Stream,
			&rs.Run, &rs.Portsrv, &rs.Sp, &rs.Camid, &rs.Record_status,
			&rs.Stream_status, &rs.Record_state, &rs.Stream_state)
		if err != nil {
			logger.LogError(s.log, err)
			return refreshStream.RefreshStreamWithNull{}, err
		}
	}

	if rs.Id == 0 {
		return refreshStream.RefreshStreamWithNull{}, fmt.Errorf("this Id does not exist")
	}

	logger.LogDebug(s.log, constants.GetIdRespConst)
	return rs, nil
}

func (s refreshStreamRepository) Insert(ctx context.Context,
	nameColumns string, values string) error {

	valSlice := strings.Split(values, ", ")
	if len(strings.Split(nameColumns, ", ")) != len(valSlice) {
		return fmt.Errorf(constants.InsertRespErrConst)
	}

	_, err := s.db.ExecContext(ctx, fmt.Sprintf(constants.InsertQueryConst, nameColumns, values))
	if err != nil {
		return err
	}

	logger.LogDebug(s.log, "success insertn")
	return nil
}

// изменяет одну ячейку
func (s refreshStreamRepository) Update(ctx context.Context, setCol string,
	setValue, whereValue interface{}) error {

	if whereValue == 0 {
		return fmt.Errorf("this Id does not exist")
	}

	_, err := s.db.ExecContext(ctx, fmt.Sprintf(constants.UpdateQueryConst, setCol), whereValue, setValue)
	if err != nil {
		return err
	}

	logger.LogDebug(s.log, fmt.Sprintf(constants.UpdateContextRespConst, setCol, setValue))
	return nil
}

// удаление записи по значению value столбца column
func (s refreshStreamRepository) Delete(ctx context.Context, value int) error {

	_, err := s.db.ExecContext(ctx, constants.DeleteContextQueryConst, value)
	if err != nil {
		return err
	}

	logger.LogDebug(s.log, fmt.Sprintf(constants.DeleteContextRespConst))
	return nil
}
