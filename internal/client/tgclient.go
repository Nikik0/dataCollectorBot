package client

import (
	"github.com/Nikik0/dataCollectorBot/internal/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

type TgClient struct {
	client *tgbotapi.BotAPI
}

func New(token *string) TgClient {
	bot, err := tgbotapi.NewBotAPI(*token)
	if err != nil {
		logger.Fatal("Failed to start bot. Check if the token is valid.")
	}
	return TgClient{bot}
}

func (tg *TgClient) SendMessage(msg string, userId int64) error {
	message := tgbotapi.NewMessage(userId, msg)
	_, err := tg.client.Send(message)
	if err != nil {
		return errors.Wrap(err, "Exception when sending msg to user")
	}
	return nil
}

func (tg *TgClient) ListenUpdates() {
	uConfig := tgbotapi.NewUpdate(0)
	uConfig.Timeout = 60

	upds := tg.client.GetUpdatesChan(uConfig)
	logger.Info("App started listening for updates.")
	for upd := range upds {
		processUpdate(upd, tg)
	}
}

func processUpdate(update tgbotapi.Update, tg *TgClient) {
}
