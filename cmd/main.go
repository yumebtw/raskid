package main

import (
	"github.com/yumebtw/raskid/internal/bot"
	"github.com/yumebtw/raskid/internal/config"
	"github.com/yumebtw/raskid/internal/db"
	"log"
)

func main() {
	cfg := config.MustLoad()
	db.ConnectDB(cfg.Database)
	_, err := bot.NewTelegramBot(cfg.Bot.Token)
	if err != nil {
		log.Fatal(err)
	}
}
