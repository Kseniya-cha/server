package refreshstream

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Kseniya-cha/server/constants"
	dblog "github.com/Kseniya-cha/server/pkg/DBLog"
)

func SelectContextRS(ctx context.Context, db dblog.DBLog) ([]byte, error) {

	rows, err := db.Db.QueryContext(ctx, constants.SelectContextRSQueryConst)
	if err != nil {
		return nil, err
	}

	db.LogIFAction(err, constants.SelectContextRSRespConst)
	defer rows.Close()

	refreshStreamArr := []RefreshStreamWithNull{}

	var json_data []byte
	// var datas [][]byte
	for rows.Next() {
		rs := RefreshStreamWithNull{}

		err := rows.Scan(&rs.Id, &rs.Auth, &rs.Ip, &rs.Stream,
			&rs.Run, &rs.Portsrv, &rs.Sp, &rs.Camid, &rs.Record_status,
			&rs.Stream_status, &rs.Record_state, &rs.Stream_state)
		if err != nil {
			return nil, err
		}
		refreshStreamArr = append(refreshStreamArr, rs)
	}
	json_data, err = json.Marshal(refreshStreamArr)
	db.LogPrintFat(err)

	return json_data, nil
}

func GetIdContextRS(ctx context.Context, db dblog.DBLog,
	value interface{}) ([]byte, error) {

	rows, err := db.Db.QueryContext(ctx, fmt.Sprintf(constants.GetIdQueryConst, "id"), value)
	if err != nil {
		return nil, err
	}
	db.LogIFAction(err, constants.GetIdContextActConst)
	defer rows.Close()

	refreshStreamArr := []RefreshStreamWithNull{}
	var json_data []byte

	for rows.Next() {
		rs := RefreshStreamWithNull{}

		err := rows.Scan(&rs.Id, &rs.Auth, &rs.Ip, &rs.Stream,
			&rs.Run, &rs.Portsrv, &rs.Sp, &rs.Camid, &rs.Record_status,
			&rs.Stream_status, &rs.Record_state, &rs.Stream_state)
		if err != nil {
			return nil, err
		}
		refreshStreamArr = append(refreshStreamArr, rs)
		db.LogPrintFat(err)
		json_data, err = json.Marshal(rs)
	}
	db.LogI(constants.GetIdContextRespConst)
	return json_data, nil
}

// nameColumns записывать в виде `"col1", "col2"`,
// values - в виде `'val1', 'val2'`
func InsertContext(ctx context.Context, db dblog.DBLog,
	nameColumns string, values string) error {

	valSlice := strings.Split(values, ", ")
	if len(strings.Split(nameColumns, ", ")) != len(valSlice) {
		db.LogF(constants.InsertRespErrConst)
		return fmt.Errorf(constants.InsertRespErrConst)
	}

	_, err := db.Db.ExecContext(ctx, fmt.Sprintf(constants.InsertQueryConst, nameColumns, values))
	if err != nil {
		return err
	}
	return nil
}

// изменяет одну ячейку
func UpdateContext(ctx context.Context, db dblog.DBLog, setCol string,
	setValue, whereValue interface{}) error {

	_, err := db.Db.ExecContext(ctx, fmt.Sprintf(constants.UpdateQueryConst, setCol), whereValue, setValue)
	if err != nil {
		return err
	}
	db.LogI(fmt.Sprintf(constants.UpdateContextRespConst, setCol, setValue))
	return nil
}

// удаление записи по значению value столбца column
func DeleteContext(ctx context.Context, db dblog.DBLog, value int) error {

	_, err := db.Db.ExecContext(ctx, constants.DeleteContextQueryConst, value)
	if err != nil {
		return err
	}
	db.LogI(fmt.Sprintf(constants.DeleteContextRespConst))
	return nil
}
