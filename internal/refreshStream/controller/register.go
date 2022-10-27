package controller

import (
	"context"
	"database/sql"

	"github.com/Kseniya-cha/server/constants"
	refreshStream "github.com/Kseniya-cha/server/internal/refreshStream"
	"github.com/Kseniya-cha/server/pkg/logger"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func RegisterRouter(router *mux.Router, useCase refreshStream.RefreshStreamUseCase, db *sql.DB, log *logrus.Logger) {
	ctx := context.Background()
	h := NewRefreshStreamHandler(useCase, db, log)

	hfSelect := h.GetAllHF(ctx)
	router.HandleFunc(constants.URLApiConst, hfSelect).Methods("GET")

	// http://localhost:3333/api/3/
	hfGetId := h.GetIdHF(ctx)
	router.HandleFunc(constants.URLGetDelIdConst, hfGetId).Methods("GET")

	// http://localhost:3333/api/3/
	hfDeleteId := h.DeleteIdHF(ctx)
	router.HandleFunc(constants.URLGetDelIdConst, hfDeleteId).Methods("DELETE")

	hfPostJS := h.PostHFJSON(ctx)
	router.HandleFunc(constants.URLApiConst,
		hfPostJS).Methods("POST")

	hfPutJS := h.PutHFJSON(ctx)
	router.HandleFunc(constants.URLApiConst,
		hfPutJS).Methods("PUT")

	hfPatchJS := h.PatchHFJSON(ctx)
	router.HandleFunc(constants.URLApiConst,
		hfPatchJS).Methods("PATCH")

	logger.LogDebug(h.log, "handlers registered")
}
