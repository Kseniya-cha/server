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
	PostHFRespOkConst   = "Success post request. Status code: %v"
	PutHFRespOkConst    = "Success put request. Status code: %v"
	PatchHFRespOkConst  = "Success patch request. Status code: %v"

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
	SELECT_COL_FROM_TBL               = "SELECT %s FROM %s"
	SELECT_COL_FROM_TBL_WHERE_CND     = "SELECT %s FROM %s WHERE %s"
	INSERT_INTO_TBL_VALUES_VAL        = "INSERT INTO %s(%s) VALUES (%s) ON CONFLICT DO NOTHING"
	UPDATE_TBL_SET_VAL_WHERE_CND      = `UPDATE public."refresh_stream" SET %v WHERE %v`
	DELETE_FROM_TBL_WHERE_CND         = "DELETE FROM %s WHERE %s"
	DELETE_CASCADE_FROM_TBL_WHERE_CND = "DELETE CASCADE FROM %s WHERE %s"
)

// msg
const (
	GetRespErrConst        = "cannot get: %v"
	DeleteRespConst        = "Success delete"
	DeleteRespErrConst     = "cannot delete: %v"
	InsertRespOkConst      = "Success insert"
	InsertRespErrConst     = "cannot insert: %v"
	UpdateHFAllValuesConst = `'%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v'`
	UpdateRespErrConst     = "cannot update: %v"

	DBRespConst         = "Received response from the database"
	IDDoesNotExistConst = "this Id does not exist"
)
