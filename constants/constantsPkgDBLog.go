package constants

// postgres
const OpenDBConst = "host=%s port=%s user=%s password=%s dbname=%s"
const OpenDBRespConst = "open database %s"
const OpenDBPingRespConst = "connection to database %s"
const CloseDBRespConst = "close database %s"
const DriverName = "postgres"

// gracefulShutdown
const SigConst = "got signal: %v, exiting"

// ping
const ConnOkConst = "connectoion is ok"
const ConnErrConst = "connection was destroy! waiting for connection..."
