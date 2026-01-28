package db

import (
	"category-crud/config"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Configure(config config.Template) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.DB.ConnectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if config.App.Debug {
		log.Default().Println("Connected to database")
	}

	db.SetMaxOpenConns(config.DB.MaxOpenConns)
	db.SetMaxIdleConns(config.DB.MaxIdleConns)

	return db, nil

}
