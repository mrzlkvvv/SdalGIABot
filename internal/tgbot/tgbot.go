package tgbot

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func NewTelegramBot() *bot.Bot {
	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
		bot.WithHTTPClient(
			time.Second*5,
			newHTTPClient(getProxyURL()),
		),
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		panic("environment variable 'TELEGRAM_BOT_TOKEN' is not set")
	}

	bot, err := bot.New(token, opts...)
	if err != nil {
		panic(err)
	}

	return bot
}

func newHTTPClient(proxyURL *url.URL) *http.Client {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	if proxyURL != nil {
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	return client
}

func getProxyURL() *url.URL {
	proxyRawURL := os.Getenv("PROXY_URL")
	if proxyRawURL == "" {
		return nil
	}

	proxyURL, err := url.Parse(proxyRawURL)
	if err != nil {
		panic(err)
	}

	err = checkProxyAvailability(proxyURL)
	if err != nil {
		panic(fmt.Errorf("proxy %s is unavailable: %w", proxyURL.String(), err))
	}

	return proxyURL
}

func checkProxyAvailability(proxyURL *url.URL) error {
	conn, err := net.DialTimeout("tcp", proxyURL.Host, 5*time.Second)
	if err != nil {
		return err
	}

	_ = conn.Close()
	return nil
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, _ = b.CopyMessage(ctx, &bot.CopyMessageParams{
		ChatID:     update.Message.Chat.ID,
		FromChatID: update.Message.Chat.ID,
		MessageID:  update.Message.ID,
	})
}
