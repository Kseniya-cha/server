package refreshStream

// name columns
const (
	IdConst           = `"id"`
	AuthConst         = `"auth"`
	IpConst           = `"ip"`
	StreamConst       = `"stream"`
	RunConst          = `"run"`
	PortsrvConst      = `"portsrv"`
	SpConst           = `"sp"`
	CamidConst        = `"camid"`
	RecordStatusConst = `"record_status"`
	StreamStatusConst = `"stream_status"`
	RecordStateConst  = `"record_state"`
	StreamStateConst  = `"stream_state"`
)

// controllers

// handleFuncs
const (
	GetHFRespOkConst    = "Success select all rows. Status code: %v"
	GetIdHFRespOkConst  = "Success select by Id = %v. Status code: %v"
	DeleteHFRespOkConst = "Success delete by Id = %v. Status code: %v"

	PostHFAllColsConst   = `"auth","ip","stream","run","portsrv","sp","camid","record_status","stream_status","record_state","stream_state"`
	PostHFAllValuesConst = `'%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v'`
	PostHFRespOkConst    = "Success post request. Status code: %v"

	PutHFRespOkConst   = "Success put request. Status code: %v"
	PatchHFRespOkConst = "Success patch request. Status code: %v"
	PutRespErrColConst = "cannot update: %v"

	ConvertIdIntConst = "Success converse Id to int"
	DecodeJsonConst   = "Success decode json"
)

// register
const (
	URLApiConst   = "/api/"
	URLApiIdConst = "/api/{ID}/"

	RegisteredHandlerOkConst = "Handlers registered!"
)

// usecase

const (
	ConvertIDErrConst = "cannot convert Id to int"
)

// repository

// queries
const (
	GetQueryConst    = `select * from public."refresh_stream"`
	GetIDQueryConst  = `select * from public."refresh_stream" where %s=$1`
	DeleteQueryConst = `delete from public."refresh_stream" where "id" = $1`
	InsertQueryConst = `insert into public."refresh_stream"(%s) values(%s)`
	UpdateQueryConst = `update public."refresh_stream" set %s=$2 where "id"=$1`
)
const (
	INSERT_INTO_TBL_VALUES_VAL        = "INSERT INTO %s VALUES %s ON CONFLICT DO NOTHING"
	SELECT_COL_FROM_TBL_WHERE_CND     = "SELECT %s FROM %s WHERE %s"
	SELECT_COL_FROM_TBL               = "SELECT %s FROM %s"
	UPDATE_TBL_SET_VAL_WHERE_CND      = `UPDATE public."refresh_stream" SET %v WHERE %v`
	DELETE_FROM_TBL_WHERE_CND         = "DELETE FROM %s WHERE %s"
	DELETE_CASCADE_FROM_TBL_WHERE_CND = "DELETE CASCADE FROM %s WHERE %s"
)

// msg
const (
	GetSentRespConst            = "Sent query to database"
	GetRespErrConst             = "cannot get: %v"
	DeleteRespConst             = "Success delete"
	DeleteRespErrConst          = "cannot delete: %v"
	InsertRespErrCountColsConst = "cannot insert: more columns than values"
	InsertRespErrConst          = "cannot insert: %v"
	InsertRespOkConst           = "Success insert"
	UpdateRespOkConst           = "Success update: %v = %v"
	UpdateRespErrConst          = "cannot update: %v"

	DBRespConst         = "Received response from the database"
	IDDoesNotExistConst = "this Id does not exist"
)
