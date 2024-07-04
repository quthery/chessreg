package database

import (
	"chessreg/internal/attrs"
	"database/sql"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func InitDB(path string) *Storage {
	db, _ := sql.Open("sqlite3", path)

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		wins INTEGER DEFAULT 0,
		loses INTEGER DEFAULT 0,
		winrate INTEGER DEFAULT 0,
		matches INTEGER DEFAULT 0
	)
`)
	if err != nil {
		slog.Error("DATABASE ERROR:", attrs.Err(err))
		return nil
	}

	_, err = stmt.Exec()

	if err != nil {
		slog.Error("DATABASE ERROR:", attrs.Err(err))
	}

	return &Storage{db: db}

}
