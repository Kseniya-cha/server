package constants

// queries
const GetQueryConst = `select * from public."refresh_stream"`
const GetSentRespConst = "sent query to database"
const GetIdQueryConst = `select * from public."refresh_stream" where %s=$1`
const GetIdRespConst = "database response accepted"
const InsertQueryConst = `insert into public."refresh_stream"(%s) values(%s)`
const InsertRespErrConst = "more columns than values! insert break"
const UpdateQueryConst = `update public."refresh_stream"set %s=$2 where "id"=$1`
const UpdateContextRespConst = "success update: %v = %v"
const DeleteContextQueryConst = `delete from public."refresh_stream" where "id" = $1`
const DeleteContextRespConst = "success delete"
