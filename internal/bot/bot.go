package bot

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	m "github.com/yumebtw/raskid/internal/models"
	"log"
	"strings"

	"github.com/yumebtw/raskid/internal/db"
)

var pendingNades = map[int64]*m.Nade{}

var userStates = make(map[int64]string)

func NewTelegramBot(s *db.Storage, token string) (*bot.Bot, error) {
	b, err := bot.New(token, bot.WithDefaultHandler(DefaultHandler(s)))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Bot initialized")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	b.Start(ctx)

	return b, nil
}

func DefaultHandler(s *db.Storage) func(ctx context.Context, b *bot.Bot, update *models.Update) {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update == nil {
			return
		}

		chatID := update.Message.Chat.ID
		text := update.Message.Text

		var response string

		if text == "/start" {
			response = "Welcome to raskid bot! Use /help to see all commands."
		} else if text == "/help" {
			response = `Type /nades <map> <team> <vector> <type> <usage> <position> to get a video for this nade
<map> - is the map for which you need your nade (e.g. mirage, inferno, ancient, etc)
<team> - T or CT?
<vector> - what is the part of the map where you are going to use your nade (A, Mid or B?)
<class> - nade class (e.g. smoke, flash, molly and he)
<usage> - the purpose of the nade, e.g. defence(for CTs), execute(for Ts), retake, insta(for smokes), assist(support flash), selfpop(flashes only)
<position> - the position where you want to use your nade
example of command: "/nades mirage t mid smoke insta window"`
		} else if strings.HasPrefix(text, "/nades") {
			HandleGetNades(s, ctx, b, update)
		} else if strings.HasPrefix(text, "/addnade") {
			HandleAddNade(ctx, b, update)
		} else if userStates[chatID] == "awaiting_url" {
			HandleNadeLink(s, ctx, b, update)
		} else {
			response = "Unknown command. Use /help to see all commands."
		}

		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   response,
		})
		if err != nil {
			log.Println(err)
		}
	}
}

func HandleGetNades(s *db.Storage, ctx context.Context, b *bot.Bot, update *models.Update) {
	parts := strings.Split(update.Message.Text, " ")
	if len(parts) != 7 {
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Invalid command. We did not implement incomplete /nades command support yet",
		})
		if err != nil {
			log.Println(err)
		}

		return
	}

	links, err := s.GetNades(parts[1], parts[2], parts[3], parts[4], parts[5], parts[6])
	if err != nil {
		log.Println(err)
	}

	fmt.Println(links[0])

	var response strings.Builder
	response.WriteString(fmt.Sprintf("🔥 Nades for **%s** (%s):\n\n", parts[0], parts[1]))
	for _, link := range links {
		fmt.Println("link: ", link)
		response.WriteString(fmt.Sprintf(
			"▶️ [Watch Video](%s)\n\n",
			link,
		))
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   response.String(),
	})
	if err != nil {
		log.Println(err)
	}
}

func HandleAddNade(ctx context.Context, b *bot.Bot, update *models.Update) {
	userID := update.Message.Chat.ID

	parts := strings.Split(update.Message.Text, " ")
	if len(parts) < 7 {
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: userID,
			Text:   "We are sorry, we did not implement incomplete /nades command feature yet.",
		})

		if err != nil {
			log.Println(err)
		}

		return
	}

	userStates[userID] = "awaiting_url"
	pendingNades[userID] = &m.Nade{
		Map:      parts[1],
		Team:     parts[2],
		Vector:   parts[3],
		Class:    parts[4],
		Usage:    parts[5],
		Position: parts[6],
	}

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: userID,
		Text:   "Please send me the URL you want to add:",
	})

	if err != nil {
		log.Println(err)
	}

}

func HandleNadeLink(s *db.Storage, ctx context.Context, b *bot.Bot, update *models.Update) {
	userID := update.Message.Chat.ID
	text := update.Message.Text

	nade := pendingNades[userID]

	_, err := s.AddNade(nade.Map, nade.Team, nade.Vector, nade.Class, nade.Usage, nade.Position, text)
	if err != nil {
		log.Println(err)
	}

	delete(userStates, userID)
	delete(pendingNades, userID)

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: userID,
		Text:   "The nade added successfully!.",
	})
}
