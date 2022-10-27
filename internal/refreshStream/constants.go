package refreshStream

// controllers

// handleFuncs
const (
	GetHFRespOkConst    = "Success select all rows!"
	GetIdHFRespOkConst  = "Success select by Id = %d!"
	DeleteHFRespOkConst = "Success delete by Id = %d!"

	PostHFAllColsConst   = `"auth","ip","stream","run","portsrv","sp","camid","record_status","stream_status","record_state","stream_state"`
	PostHFAllValuesConst = `'%s','%s','%s','%s','%s','%s','%s','%v','%v','%v','%v'`
	PostHFRespOkConst    = "Success post request"

	ConvertIdIntConst = "Success converse Id to int"
	DecodeJsonConst   = "Success decode json"
)

// register
const (
	URLApiConst   = "/api/"
	URLApiIdConst = "/api/{ID}/"

	RegisteredHandlerOkConst = "Handlers registered!"
)

// repository

// queries
const (
	GetQueryConst    = `select * from public."refresh_stream"`
	GetIDQueryConst  = `select * from public."refresh_stream" where %s=$1`
	DeleteQueryConst = `delete from public."refresh_stream" where "id" = $1`
	InsertQueryConst = `insert into public."refresh_stream"(%s) values(%s)`
	UpdateQueryConst = `update public."refresh_stream"set %s=$2 where "id"=$1`
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
