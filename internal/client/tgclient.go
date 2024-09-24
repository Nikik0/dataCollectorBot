package client

import (
	"fmt"
	"github.com/Nikik0/dataCollectorBot/internal/cache"
	"github.com/Nikik0/dataCollectorBot/internal/logger"
	"github.com/Nikik0/dataCollectorBot/internal/model"
	"github.com/Nikik0/dataCollectorBot/internal/statemachine"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

type TgClient struct {
	client         *tgbotapi.BotAPI
	tgStateMachine *statemachine.StateMachine
	lru            *cache.LRU
}

func New(token *string) TgClient {
	bot, err := tgbotapi.NewBotAPI(*token)
	if err != nil {
		logger.Fatal("Failed to start bot. Check if the token is valid.")
	}
	sm := statemachine.NewStateMachine()
	lru := cache.NewLRU(1000)
	return TgClient{bot, sm, lru}
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
	if update.Message != nil {
		logger.Info(fmt.Sprintf("[%s][%v] %s", update.Message.From.UserName, update.Message.From.ID, update.Message.Text))
		processMessage(model.Message{
			Text:            update.Message.Text,
			UserID:          update.Message.From.ID,
			UserName:        update.Message.From.UserName,
			UserDisplayName: strings.TrimSpace(update.Message.From.FirstName + " " + update.Message.From.LastName),
			IsCallback:      false,
			CallbackMsgID:   "",
		}, tg)
	}
}

func processMessage(msg model.Message, tg *TgClient) {
	u := tg.lru.Get(strconv.FormatInt(msg.UserID, 10)) //todo get user from lru cache
	if u == nil {
		u = model.NewUser(msg.UserID)
	}
	if msg.IsCallback == false {
		if u.IsMistakeCorrectionNeeded() {
			tg.tgStateMachine.SendEvent(statemachine.ConfirmationNeeded, &u, &msg)
		} else {
			tg.tgStateMachine.SendEvent(statemachine.Default, &u, &msg)
		}
	} else {
		switch msg.CallbackMsgID { //todo more cases
		case "/changeName":
			{
				u.SetMistakeCorrectionNeeded(true)
				tg.tgStateMachine.SendEvent(statemachine.NameNeeded, &u, &msg)
			}
		case "/changeSurname":
			{
				u.SetMistakeCorrectionNeeded(true)
				tg.tgStateMachine.SendEvent(statemachine.SurnameNeeded, &u, &msg)
			}
		case "/changeBirthdate":
			{
				u.SetMistakeCorrectionNeeded(true)
				tg.tgStateMachine.SendEvent(statemachine.BirthdateNeeded, &u, &msg)
			}
		case "/changeEmail":
			{
				u.SetMistakeCorrectionNeeded(true)
				tg.tgStateMachine.SendEvent(statemachine.EmailNeeded, &u, &msg)
			}
		case "/confirmed":
			{
				u.SetMistakeCorrectionNeeded(false)
				tg.tgStateMachine.SendEvent(statemachine.FinishNeeded, &u, &msg)
			}
		}
	}
}
