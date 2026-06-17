package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/mrzlkvvv/SdalGIABot/internal/tgbot"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	tgbot := tgbot.NewTelegramBot()

	tgbot.Start(ctx)
}
