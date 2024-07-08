package sqlite

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	db *sqlx.DB
}

func New() (*Repository, error) {
	db, err := sqlx.Open("sqlite3", "./main.db")
	if err != nil {
		return &Repository{}, err
	}

	err = db.Ping()
	if err != nil {
		return &Repository{}, err
	}

	newDB := &Repository{
		db: db,
	}
	err = newDB.runMigration()
	if err != nil {
		return &Repository{}, err
	}
	fmt.Println("db connected")
	return newDB, nil
}

func (repo *Repository) Close() {
	repo.db.Close()
}

func (repo *Repository) runMigration() error {
	_, err := repo.db.Exec(`
  CREATE TABLE IF NOT EXISTS todos (
    id TEXT PRIMARY KEY,
    title TEXT,
    is_completed TEXT,
    user_id TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(id)
  )
	`)
	if err != nil {
		return err
	}

	// Создание таблицы languages
	_, err = repo.db.Exec(`
  CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    email TEXT UNIQUE,
    name TEXT,
    password TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  )
	`)
	if err != nil {
		return err
	}

	return nil
}
