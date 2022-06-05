package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	config *Config
	db     *pgxpool.Pool
}

func New(cfg *Config) *Storage {
	return &Storage{
		config: cfg,
	}
}

func (s *Storage) Open() error {
	db, err := pgxpool.Connect(context.Background(), s.config.DBURL)
	if err != nil {
		log.Print("Trouble with opening pgx connection")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		log.Println("Trouble pinging pgx conn")
		return err
	}

	s.db = db
	return nil
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) WriteUser() error {
	return nil
}

func (s *Storage) ReadUser() *User {
	return &User{}
}
