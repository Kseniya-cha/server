package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

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

	logger.LogDebug(s.log, refreshStream.GetSentRespConst)
	rows, err := s.db.QueryContext(ctx, refreshStream.GetQueryConst)
	if err != nil {
		logger.LogError(s.log, refreshStream.GetRespErrConst)
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
	logger.LogDebug(s.log, refreshStream.DBRespConst)
	return refreshStreamArr, nil
}

func (s refreshStreamRepository) GetId(ctx context.Context,
	value interface{}) (refreshStream.RefreshStreamWithNull, error) {

	rows, err := s.db.QueryContext(ctx, fmt.Sprintf(refreshStream.GetIDQueryConst, "id"), value)
	if err != nil {
		logger.LogError(s.log, refreshStream.GetRespErrConst)
		return refreshStream.RefreshStreamWithNull{}, err
	}
	logger.LogDebug(s.log, refreshStream.GetSentRespConst)

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
		return refreshStream.RefreshStreamWithNull{}, fmt.Errorf(refreshStream.IDDoesNotExistConst)
	}

	logger.LogDebug(s.log, refreshStream.DBRespConst)
	return rs, nil
}

// удаление записи по значению value столбца column
func (s refreshStreamRepository) Delete(ctx context.Context, value int) error {

	_, err := s.db.ExecContext(ctx, refreshStream.DeleteQueryConst, value)
	if err != nil {
		return fmt.Errorf(refreshStream.DeleteRespErrConst, err)
	}

	logger.LogDebug(s.log, fmt.Sprintf(refreshStream.DeleteRespConst))
	return nil
}

func (s refreshStreamRepository) Insert(ctx context.Context,
	nameColumns string, values string) error {

	valSlice := strings.Split(values, ", ")
	if len(strings.Split(nameColumns, ", ")) != len(valSlice) {
		return fmt.Errorf(refreshStream.InsertRespErrCountColsConst)
	}

	_, err := s.db.ExecContext(ctx, fmt.Sprintf(refreshStream.InsertQueryConst, nameColumns, values))
	if err != nil {
		return fmt.Errorf(refreshStream.InsertRespErrConst, err)
	}

	logger.LogDebug(s.log, refreshStream.InsertRespOkConst)
	return nil
}

// изменяет одну ячейку
func (s refreshStreamRepository) Update(ctx context.Context, setCol string,
	setValue, whereValue interface{}) error {

	if whereValue == nil {
		return fmt.Errorf(refreshStream.IDDoesNotExistConst)
	}

	_, err := s.db.ExecContext(ctx, fmt.Sprintf(refreshStream.UpdateQueryConst, setCol), whereValue, setValue)
	if err != nil {
		return fmt.Errorf(refreshStream.UpdateRespErrConst, err)
	}

	logger.LogDebug(s.log, fmt.Sprintf(refreshStream.UpdateRespOkConst, setCol, setValue))
	return nil
}
