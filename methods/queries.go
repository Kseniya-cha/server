package methods

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Kseniya-cha/server/logger"
	"github.com/Kseniya-cha/server/model"
	"github.com/sirupsen/logrus"
)

func SelectContextRS(ctx context.Context, db *sql.DB,
	log *logrus.Logger) []model.RefreshStreamWithNull {

	rows, err := db.QueryContext(ctx, `select * from public."refresh_stream"`)
	logger.LogIFAction(log, err, "sent query to database")
	defer rows.Close()

	refreshStreamArr := []model.RefreshStreamWithNull{}

	for rows.Next() {
		rs := model.RefreshStreamWithNull{}

		err := rows.Scan(&rs.Id, &rs.Auth, &rs.Ip, &rs.Stream,
			&rs.Run, &rs.Portsrv, &rs.Sp, &rs.Camid, &rs.Record_status,
			&rs.Stream_status, &rs.Record_state, &rs.Stream_state)
		if err != nil {
			log.Warn(err)
			continue
		}

		refreshStreamArr = append(refreshStreamArr, rs)

	}
	log.Infof("database response accepted")
	return refreshStreamArr
}

func GetIdContextRS(ctx context.Context, db *sql.DB, log *logrus.Logger,
	value interface{}) []model.RefreshStreamWithNull {

	rows, err := db.QueryContext(ctx, fmt.Sprintf(`select * from public."refresh_stream" 
	where %s=$1`, "id"), value)
	logger.LogIFAction(log, err, "sent query to database")
	defer rows.Close()

	refreshStreamArr := []model.RefreshStreamWithNull{}

	for rows.Next() {
		rs := model.RefreshStreamWithNull{}

		err := rows.Scan(&rs.Id, &rs.Auth, &rs.Ip, &rs.Stream,
			&rs.Run, &rs.Portsrv, &rs.Sp, &rs.Camid, &rs.Record_status,
			&rs.Stream_status, &rs.Record_state, &rs.Stream_state)
		if err != nil {
			log.Warn(err)
			continue
		}

		refreshStreamArr = append(refreshStreamArr, rs)

	}
	log.Infof("database response accepted")
	return refreshStreamArr
}

// nameColumns записывать в виде `"col1", "col2"`,
// values - в виде `'val1', 'val2'`
func InsertContext(ctx context.Context, db *sql.DB, log *logrus.Logger,
	nameColumns string, values string) {

	valSlice := strings.Split(values, ", ")
	if len(strings.Split(nameColumns, ", ")) != len(valSlice) {
		log.Warn("more columns than values! insert break")
		return
	}

	_, errIns := db.ExecContext(ctx, fmt.Sprintf(`insert into public."refresh_stream"(%s)
	values(%s)`, nameColumns, values))
	logger.LogIFAction(log, errIns, "insert")
}

// изменяет одну ячейку
func UpdateContext(ctx context.Context, db *sql.DB, log *logrus.Logger,
	setCol string, setValue, whereValue interface{}) error {

	_, err := db.ExecContext(ctx, fmt.Sprintf(`update public."refresh_stream"
	set %s=$2 where "id"=$1`, setCol), whereValue, setValue)
	if err != nil {
		return err
	}
	return nil
}

// удаление записи по значению value столбца column
func DeleteContext(ctx context.Context, db *sql.DB, log *logrus.Logger,
	column string, value interface{}) error {

	res, err := db.ExecContext(ctx, fmt.Sprintf(`delete from public."refresh_stream"
	where %s = $1`, column), value)
	logger.LogI(log, fmt.Sprintf("success delete: res = %v", res))
	if err != nil {
		return err
	}
	return nil
}
