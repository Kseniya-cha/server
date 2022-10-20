package methods

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Kseniya-cha/server/model"
)

func SelectContextRS(ctx context.Context, db model.DBLog) ([]byte, error) {

	rows, err := db.Db.QueryContext(ctx, `select * from public."refresh_stream"`)
	if err != nil {
		return nil, err
	}
	db.LogIFAction(err, "sent query to database")
	defer rows.Close()

	refreshStreamArr := []model.RefreshStreamWithNull{}

	var json_data []byte
	// var datas [][]byte
	for rows.Next() {
		rs := model.RefreshStreamWithNull{}

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

func GetIdContextRS(ctx context.Context, db model.DBLog,
	value interface{}) ([]byte, error) {

	rows, err := db.Db.QueryContext(ctx, fmt.Sprintf(`select * from public."refresh_stream" 
	where %s=$1`, "id"), value)
	if err != nil {
		return nil, err
	}
	db.LogIFAction(err, "sent query to database")
	defer rows.Close()

	refreshStreamArr := []model.RefreshStreamWithNull{}
	var json_data []byte

	for rows.Next() {
		rs := model.RefreshStreamWithNull{}

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
	db.LogI("database response accepted")
	return json_data, nil
}

// nameColumns записывать в виде `"col1", "col2"`,
// values - в виде `'val1', 'val2'`
func InsertContext(ctx context.Context, db model.DBLog,
	nameColumns string, values string) error {

	valSlice := strings.Split(values, ", ")
	if len(strings.Split(nameColumns, ", ")) != len(valSlice) {
		db.LogF("more columns than values! insert break")
		return fmt.Errorf("more columns than values! insert break")
	}

	_, err := db.Db.ExecContext(ctx, fmt.Sprintf(`insert into public."refresh_stream"(%s)
	values(%s)`, nameColumns, values))
	if err != nil {
		return err
	}
	return nil
}

// изменяет одну ячейку
func UpdateContext(ctx context.Context, db model.DBLog, setCol string,
	setValue, whereValue interface{}) error {

	_, err := db.Db.ExecContext(ctx, fmt.Sprintf(`update public."refresh_stream"
	set %s=$2 where "id"=$1`, setCol), whereValue, setValue)
	if err != nil {
		return err
	}
	db.LogI(fmt.Sprintf("success update: %v = %v", setCol, setValue))
	return nil
}

// удаление записи по значению value столбца column
func DeleteContext(ctx context.Context, db model.DBLog, value int) error {

	_, err := db.Db.ExecContext(ctx, `delete from public."refresh_stream"
	where "id" = $1`, value)
	if err != nil {
		return err
	}
	db.LogI(fmt.Sprintf("success delete"))
	return nil
}
