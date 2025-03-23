package bot

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var Bot *bot.Bot

func NewTelegramBot(token string) (*bot.Bot, error) {
	bot, err := bot.New(token, bot.WithDefaultHandler(DefaultHandler))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Bot initialized")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bot.Start(ctx)

	return bot, nil
}

func DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update == nil {
		return
	}

	chatID := update.Message.Chat.ID
	var text string

	switch text {
	case "/start":
		text = "Welcome to raskid bot! Use /help to see all commands."
	case "/help":
		text = `Type /nades <map> <team> <vector> <type> <usage> <position> to get a video for this nade
<map> - is the map for which you need your nade (e.g. mirage, inferno, ancient, etc)
<team> - T or CT?
<vector> - what is the part of the map where you are going to use your nade (A, Mid or B?)
<type> - nade type (e.g. smoke, flash, molly and he)
<usage> - the purpose of the nade, e.g. defence(for CTs), execute(for Ts), retake, insta(for smokes), assist(support flash), selfpop(flashes only)
<position> - the position where you want to use your nade
example of command: "/nades mirage t mid smoke insta window"
`
	default:
		text = "Unknown command. Use /help to see all commands."
	}

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   text,
	})
	if err != nil {
		log.Println(err)
	}
}
