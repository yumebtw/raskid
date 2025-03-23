package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/yumebtw/raskid/internal/config"
	"log"
)

type DB *sql.DB

func ConnectDB(cfg config.Database) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")
}
