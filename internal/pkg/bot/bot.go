package bot

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"

	telegramPkg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/bot/handlers"
)

type IBot interface {
	Run() error
}

func New(apiKey string, pool *pgxpool.Pool) IBot {
	bot, err := telegramPkg.NewBotAPI(apiKey)
	if err != nil {
		log.Panic(errors.Wrap(err, "failed to initialize bot"))
	}

	return &commander{
		bot: bot,
		h:   handlers.New(pool),
	}
}

type commander struct {
	bot *telegramPkg.BotAPI
	h   handlers.IHandler
}

func (c *commander) Run() error {
	u := telegramPkg.NewUpdate(0)
	u.Timeout = 60

	updates := c.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := telegramPkg.NewMessage(update.Message.Chat.ID, "")
		if cmd := update.Message.Command(); cmd != "" {
			if text, err := c.h.Call(cmd, update.Message.CommandArguments()); err == nil {
				if text != "" {
					msg.Text = text
				} else {
					msg.Text = "Done!"
				}
			} else {
				fmt.Println(err)
				msg.Text = err.Error()
			}
		} else {
			fmt.Printf("[%s] %s\n", update.Message.From.UserName, update.Message.Text)
			msg.Text = update.Message.Text
		}

		if _, err := c.bot.Send(msg); err != nil {
			return errors.Wrap(err, fmt.Sprintf("send message: %s", msg.Text))
		}
	}
	return nil
}
