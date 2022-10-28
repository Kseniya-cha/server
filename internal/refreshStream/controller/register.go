package controller

import (
	"context"
	"database/sql"

	refreshStream "github.com/Kseniya-cha/server/internal/refreshStream"
	"github.com/Kseniya-cha/server/pkg/logger"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func RegisterRouter(router *mux.Router, useCase refreshStream.RefreshStreamUseCase, db *sql.DB, log *logrus.Logger) {

	ctx := context.Background()
	h := NewRefreshStreamHandler(useCase, db, log)

	hfSelect := h.GetHF(ctx)
	router.HandleFunc(refreshStream.URLApiConst, hfSelect).Methods("GET")

	hfGetId := h.GetIdHF(ctx)
	router.HandleFunc(refreshStream.URLApiIdConst, hfGetId).Methods("GET")

	hfDeleteId := h.DeleteIdHF(ctx)
	router.HandleFunc(refreshStream.URLApiIdConst, hfDeleteId).Methods("DELETE")

	hfPostJS := h.PostHFJSON(ctx)
	router.HandleFunc(refreshStream.URLApiConst, hfPostJS).Methods("POST")

	hfPutJS := h.PutHFJSON(ctx)
	router.HandleFunc(refreshStream.URLApiConst, hfPutJS).Methods("PUT")

	hfPatchJS := h.PatchHFJSON(ctx)
	router.HandleFunc(refreshStream.URLApiConst, hfPatchJS).Methods("PATCH")

	logger.LogDebug(h.log, refreshStream.RegisteredHandlerOkConst)
}
