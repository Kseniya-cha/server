package constants

// queries
const SelectContextRSQueryConst = `select * from public."refresh_stream"`
const SelectContextRSRespConst = "sent query to database"
const GetIdQueryConst = `select * from public."refresh_stream" where %s=$1`
const GetIdContextActConst = "sent query to database"
const GetIdContextRespConst = "database response accepted"
const InsertQueryConst = `insert into public."refresh_stream"(%s) values(%s)`
const InsertRespErrConst = "more columns than values! insert break"
const UpdateQueryConst = `update public."refresh_stream"set %s=$2 where "id"=$1`
const UpdateContextRespConst = "success update: %v = %v"
const DeleteContextQueryConst = `delete from public."refresh_stream" where "id" = $1`
const DeleteContextRespConst = "success delete"

// handleFuncs
const PostHFAllColsConst = `"auth","ip","stream","run","portsrv","sp","camid","record_status","stream_status","record_state","stream_state"`
const PostHFAllValuesConst = `'%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s'`
const ConvIdIntConst = "converse id to int"
const GetAllHFRespOkConst = "success select all rows!"
const RootPageConst = "this is root page\n"
const GetIdHFRespOkConst = "success select by id = %d!"
const RecSttConst = "record_status"
const StrSttConst = "stream_status"
const RecStatConst = "record_state"
const StrStatConst = "streamstate"

// runServer
const ServCloseConst = "server closed"
const ServErrConst = "error listening for server: %s"

// gracefulShutdown
const SigConst = "got signal: %v, exiting"

// ping
const ConnOkConst = "connectoion is ok"
const ConnErrConst = "connection was destroy! waiting for connection..."
