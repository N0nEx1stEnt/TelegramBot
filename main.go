package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6118910710:AAF-lrqev-cGtd78Al69qbJKsIoNPUgpwro")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Autozied on accaunt %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Text {
		case "/start":
			msg.Text = "Привет! Я тестовый бот. \nВыбери нужное действие"
		case "Кнопки":
			msg.Text = "Что бы вы хотели?"
			msg.ReplyMarkup = createButtons()
		default:
			msg.Text = "Прошу прощения! Но я вас не понимаю :("
		}

		_, err := bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}
}

func createButtons() tgbotapi.ReplyKeyboardMarkup {
	buttonRow1 := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("Внесите кальян"),
		tgbotapi.NewKeyboardButton("Уволить Джуна"),
	}

	buttonRow2 := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("Послать заказчика"),
		tgbotapi.NewKeyboardButton("Открыть бизнес"),
	}

	buttons := [][]tgbotapi.KeyboardButton{buttonRow1, buttonRow2}

	replyKeyboardMarkup := tgbotapi.NewReplyKeyboard(buttons...)
	replyKeyboardMarkup.OneTimeKeyboard = true

	return replyKeyboardMarkup
}
