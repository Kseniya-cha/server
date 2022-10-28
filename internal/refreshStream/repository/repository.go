package repository

import (
	"context"
	"database/sql"
	"fmt"

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

	template := refreshStream.SELECT_COL_FROM_TBL
	chose := "*"
	tbl := `public."refresh_stream"`
	query := fmt.Sprintf(template, chose, tbl)

	rows, err := s.db.QueryContext(ctx, query)
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

	template := refreshStream.SELECT_COL_FROM_TBL_WHERE_CND
	chose := "*"
	tbl := `public."refresh_stream"`
	cnd := fmt.Sprintf("%s=%d", refreshStream.IdConst, id)
	query := fmt.Sprintf(template, chose, tbl, cnd)

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		logger.LogError(s.log, refreshStream.GetRespErrConst)
		return refreshStream.RefreshStreamWithNull{}, err
	}

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

func (s refreshStreamRepository) Delete(ctx context.Context, id interface{}) error {

	template := refreshStream.DELETE_FROM_TBL_WHERE_CND
	tbl := `public."refresh_stream"`
	cnd := fmt.Sprintf("%s=%d", refreshStream.IdConst, id)
	query := fmt.Sprintf(template, tbl, cnd)

	_, err := s.db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf(refreshStream.DeleteRespErrConst, err)
	}

	logger.LogDebug(s.log, fmt.Sprintf(refreshStream.DeleteRespConst))
	return nil
}

func (s refreshStreamRepository) Insert(ctx context.Context,
	rs *refreshStream.RefreshStreamWithNull) error {

	template := refreshStream.INSERT_INTO_TBL_VALUES_VAL
	tbl := `public."refresh_stream"`
	cols := fmt.Sprintf(refreshStream.AuthConst) + "," + fmt.Sprintf(refreshStream.IpConst) + "," +
		fmt.Sprintf(refreshStream.StreamConst) + "," + fmt.Sprintf(refreshStream.RunConst) + "," +
		fmt.Sprintf(refreshStream.PortsrvConst) + "," + fmt.Sprintf(refreshStream.SpConst) + "," +
		fmt.Sprintf(refreshStream.CamidConst) + "," + fmt.Sprintf(refreshStream.RecordStatusConst) + "," +
		fmt.Sprintf(refreshStream.StreamStatusConst) + "," + fmt.Sprintf(refreshStream.RecordStateConst) + "," +
		fmt.Sprintf(refreshStream.StreamStateConst)
	vals := fmt.Sprintf(refreshStream.UpdateHFAllValuesConst,
		rs.Auth.String, rs.Ip.String, rs.Stream.String, rs.Run.String,
		rs.Portsrv, rs.Sp.String, rs.Camid.String, rs.Record_status.Bool,
		rs.Stream_status.Bool, rs.Record_state.Bool, rs.Stream_state.Bool)

	query := fmt.Sprintf(template, tbl, cols, vals)

	_, err := s.db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf(refreshStream.InsertRespErrConst, err)
	}

	logger.LogDebug(s.log, refreshStream.InsertRespOkConst)
	return nil
}

func (s refreshStreamRepository) Update(ctx context.Context, rs *refreshStream.RefreshStreamWithNull) error {

	template := refreshStream.UPDATE_TBL_SET_VAL_WHERE_CND
	val := fmt.Sprintf("%v=CASE WHEN $1 <> '' THEN $1 ELSE '' END, ", refreshStream.AuthConst) +
		fmt.Sprintf("%v=CASE WHEN $2 <> '' THEN $2 ELSE '' END, ", refreshStream.IpConst) +
		fmt.Sprintf("%v=CASE WHEN $3 <> '' THEN $3 ELSE '' END, ", refreshStream.StreamConst) +
		fmt.Sprintf("%v=CASE WHEN $5 <> '' THEN $5 ELSE '' END, ", refreshStream.RunConst) +
		fmt.Sprintf("%s=$6, ", refreshStream.PortsrvConst) +
		fmt.Sprintf("%v=CASE WHEN $7 <> '' THEN $7 ELSE '' END, ", refreshStream.SpConst) +
		fmt.Sprintf("%v=CASE WHEN $8 <> '' THEN $8 ELSE '' END, ", refreshStream.CamidConst) +
		fmt.Sprintf("%v=CASE WHEN $9 <> false THEN $13 ELSE NULL::boolean END, ", refreshStream.RecordStatusConst) +
		fmt.Sprintf("%v=CASE WHEN $10 <> false THEN $14 ELSE NULL::boolean END, ", refreshStream.StreamStatusConst) +
		fmt.Sprintf("%v=CASE WHEN $11 <> false THEN $15 ELSE NULL::boolean END, ", refreshStream.RecordStateConst) +
		fmt.Sprintf("%v=CASE WHEN $12 <> false THEN $16 ELSE NULL::boolean END", refreshStream.StreamStateConst)
	cnd := fmt.Sprintf("%s=$4", refreshStream.IdConst)
	query := fmt.Sprintf(template, val, cnd)

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
