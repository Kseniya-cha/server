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
	id interface{}) (refreshStream.RefreshStreamWithNull, error) {

	rows, err := s.db.QueryContext(ctx, fmt.Sprintf(refreshStream.GetIDQueryConst, "id"), id)
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
func (s refreshStreamRepository) Delete(ctx context.Context, id interface{}) error {

	_, err := s.db.ExecContext(ctx, refreshStream.DeleteQueryConst, id)
	if err != nil {
		return fmt.Errorf(refreshStream.DeleteRespErrConst, err)
	}

	logger.LogDebug(s.log, fmt.Sprintf(refreshStream.DeleteRespConst))
	return nil
}

func (s refreshStreamRepository) Insert(ctx context.Context,
	rs *refreshStream.RefreshStreamWithNull) error {

	allcols := refreshStream.PostHFAllColsConst
	allvalues := fmt.Sprintf(refreshStream.PostHFAllValuesConst,
		rs.Auth.String, rs.Ip.String, rs.Stream.String, rs.Run.String,
		rs.Portsrv, rs.Sp.String, rs.Camid.String, rs.Record_status.Bool,
		rs.Stream_status.Bool, rs.Record_state.Bool, rs.Stream_state.Bool)

	// проверка, что число строк совпадает с числом значений
	valSlice := strings.Split(allvalues, ", ")
	if len(strings.Split(allcols, ", ")) != len(valSlice) {
		return fmt.Errorf(refreshStream.InsertRespErrCountColsConst)
	}

	_, err := s.db.ExecContext(ctx, fmt.Sprintf(refreshStream.InsertQueryConst, allcols, allvalues))
	if err != nil {
		return fmt.Errorf(refreshStream.InsertRespErrConst, err)
	}

	logger.LogDebug(s.log, refreshStream.InsertRespOkConst)
	return nil
}

// изменяет одну ячейку
func (s refreshStreamRepository) Update(ctx context.Context, rs *refreshStream.RefreshStreamWithNull) error {

	template := refreshStream.UPDATE_TBL_SET_VAL_WHERE_CND
	val := fmt.Sprintf("%v=CASE WHEN $1 <> '' THEN $1 ELSE '' END, ", refreshStream.AuthConst) +
		fmt.Sprintf("%v=CASE WHEN $2 <> '' THEN $2 ELSE '' END, ", refreshStream.IpConst) +
		fmt.Sprintf("%v=CASE WHEN $3 <> '' THEN $3 ELSE '' END, ", refreshStream.StreamConst) +
		fmt.Sprintf("%v=CASE WHEN $5 <> '' THEN $5 ELSE '' END, ", refreshStream.RunConst) +
		fmt.Sprintf("%s=$6, ", refreshStream.PortsrvConst) +
		fmt.Sprintf("%v=CASE WHEN $7 <> '' THEN $7 ELSE '' END, ", refreshStream.SpConst) +
		fmt.Sprintf("%v=CASE WHEN $8 <> '' THEN $8 ELSE '' END, ", refreshStream.CamidConst) +
		fmt.Sprintf("%v=CASE WHEN $9 <> false THEN $13 ELSE false END, ", refreshStream.RecordStatusConst) +
		fmt.Sprintf("%v=CASE WHEN $10 <> false THEN $14 ELSE false END, ", refreshStream.StreamStatusConst) +
		fmt.Sprintf("%v=CASE WHEN $11 <> false THEN $15 ELSE false END, ", refreshStream.RecordStateConst) +
		fmt.Sprintf("%v=CASE WHEN $12 <> false THEN $16 ELSE false END", refreshStream.StreamStateConst)
	cnd := fmt.Sprintf("%s=$4", refreshStream.IdConst)
	query := fmt.Sprintf(template, val, cnd)
	fmt.Println(query)

	_, err := s.db.ExecContext(ctx, query, rs.Auth.String, rs.Ip.String,
		rs.Stream.String, rs.Id, rs.Run.String, rs.Portsrv, rs.Sp.String,
		rs.Camid.String, rs.Record_status.Valid, rs.Stream_status.Valid,
		rs.Record_state.Valid, rs.Stream_state.Valid, rs.Record_status.Bool,
		rs.Stream_status.Bool, rs.Record_state.Bool, rs.Stream_state.Bool)

	if err != nil {
		return fmt.Errorf(refreshStream.UpdateRespErrConst, err)
	}

	return nil
}
