package server

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Server struct {
	httpserver *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	handler_for_slog := slog.NewJSONHandler(os.Stdout, nil)

	logger := slog.NewLogLogger(handler_for_slog, slog.LevelError)

	s.httpserver = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		ErrorLog:       logger,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpserver.ListenAndServe()
}
