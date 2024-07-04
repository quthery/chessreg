package main

import (
	"chessreg/internal/attrs"
	"chessreg/internal/database"
	"chessreg/internal/handlers"
	"chessreg/internal/server"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func main() {
	storage := database.InitDB("data/main.db")

	gin.SetMode(gin.DebugMode)

	handler := handlers.NewHandler(storage)

	srv := new(server.Server)

	if err := srv.Run("8000", handler.InitRouter()); err != nil {
		slog.Error("starting error", attrs.Err(err))
	}
}
