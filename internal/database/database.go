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
		age string INTEGER NOT NULL,
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

func (s *Storage) NewUser(username string, age int) (int, error) {
	stmt, err := s.db.Prepare("INSERT INTO users(username, age) VALUES(?, ?)")
	if err != nil {
		slog.Error("DATABASE ERROR", attrs.Err(err))
		return 0, err
	}
	result, err := stmt.Exec(username, age)

	if err != nil {
		slog.Error("DATABASE ERROR", attrs.Err(err))
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		slog.Error("DATABASE ERROR", attrs.Err(err))
		return 0, err
	}
	return int(id), nil

}

func (s *Storage) DropTable() int {
	query := "DROP TABLE users;"
	_, err := s.db.Exec(query)
	if err != nil {
		slog.Error("database error:", attrs.Err(err))
		return 500
	}

	return 200

}
