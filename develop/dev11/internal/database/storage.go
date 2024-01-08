package database

import (
	"dev11/internal/config"
	"dev11/internal/model"
	"dev11/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS "events" (
    "event_id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "user_id" INTEGER NOT NULL,
    "description" TEXT NOT NULL,
    "date" TEXT NOT NULL 
)`

type IStorage interface {
	CreateEvent(event model.Event) (int, error)
	UpdateEvent(event model.Event) error
	DeleteEvent(eventID int) error
	GetForDay(userID int, day utils.CustomDate) ([]model.Event, error)
	GetForRange(userID int, begin, end utils.CustomDate) ([]model.Event, error)
	Close() error
}

type Storage struct {
	db *sqlx.DB
}

func NewStorage(cfg *config.Config) (*Storage, error) {
	db, err := sqlx.Connect("sqlite3", cfg.DSN)
	if err != nil {
		return nil, err
	}

	db.MustExec(schema)
	return &Storage{db}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
