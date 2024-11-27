package server

import (
	"context"
	"net/http"
	"time"
)

type Api struct {
	server http.Server
}

func (a *Api) Run(port string, handler http.Handler) error {

	a.server = http.Server{
		Addr:           port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return a.server.ListenAndServe()
}

func (a *Api) Stop(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}
