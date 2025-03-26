package main

import (
	"github.com/yumebtw/raskid/internal/bot"
	"github.com/yumebtw/raskid/internal/config"
	"github.com/yumebtw/raskid/internal/db"
	"log"
)

func main() {
	cfg := config.MustLoad()

	storage, err := db.ConnectDB(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	_, err = bot.NewTelegramBot(storage, cfg.Bot.Token)
	if err != nil {
		log.Fatal(err)
	}
}
