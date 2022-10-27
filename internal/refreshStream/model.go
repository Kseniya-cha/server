package refreshStream

import "database/sql"

// структура таблицы refresh_stream
// sql.Null* когда возможен null в столбце
type RefreshStreamWithNull struct {
	Id            int
	Auth          sql.NullString
	Ip            sql.NullString
	Stream        sql.NullString
	Run           sql.NullString
	Portsrv       string
	Sp            sql.NullString
	Camid         sql.NullString
	Record_status sql.NullBool
	Stream_status sql.NullBool
	Record_state  sql.NullBool
	Stream_state  sql.NullBool
}
