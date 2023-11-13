package main

import (
	"gitlab.ozon.dev/vss414/hw-1/internal/config"
	"gitlab.ozon.dev/vss414/hw-1/internal/database"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/bot"
	"log"
)

func main() {
	c, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	pool := database.New()
	defer pool.Close()

	cmd := bot.New(c.TelegramApiKey, pool)

	if err := cmd.Run(); err != nil {
		log.Fatalf("failed to run bot: %s", err)
	}
}
