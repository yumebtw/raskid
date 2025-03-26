package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/yumebtw/raskid/internal/config"
	"log"
)

type Storage struct {
	db *sql.DB
}

func ConnectDB(cfg config.Database) (*Storage, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to database")

	return &Storage{db: db}, nil
}

func (s *Storage) AddNade(gameMap, team, vector, class, usage, position, link string) (int64, error) {
	const op = "storage.PostgreSQL.AddNade"

	stmt, err := s.db.Prepare("INSERT INTO nade(map, team, vector, class, usage, position, link) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(gameMap, team, vector, class, usage, position, link)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (s *Storage) GetNades(gameMap, team, vector, class, usage, position string) ([]string, error) {
	const op = "storage.GetNades"

	rows, err := s.db.Query("SELECT link FROM nade;")
	if err != nil {
		fmt.Println("AAAAAA ASHIBKAAA")
		return nil, err
	}
	defer rows.Close()

	var links []string
	for rows.Next() {
		var link string
		cols, _ := rows.Columns()
		fmt.Println(cols)
		err = rows.Scan(&link)
		links = append(links, link)
	}

	return links, nil
}
