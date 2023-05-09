package main

import (
	"TelegramBot/buttons"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("Bot-Api")
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
			msg.ReplyMarkup = buttons.CreateButtons()

		case "/help":
			msg.Text = "Выберите нужную команду"
			msg.ReplyMarkup = buttons.CreateButtons()

		case "Внесите кальян":
			msg.Text = "Вижу ты сильно заработался, что покурим сегодня?"
			msg.ReplyMarkup = buttons.CreateHookahButtons()

		case "Мне покрепче!":
			msg.Text = "Как скажете, 15 минут и он у вас!"

		case "Стандартно и побыстрее!":
			msg.Text = "Понял вас, скоро все будет готово!"

		case "Уволить Джуна":
			msg.Text = "Он что забыл принести вам кофе на завтрак? Уверены в своем решении?"
			msg.ReplyMarkup = buttons.CreateJuniorButtons()

		case "Да, надо уволить!":
			msg.Text = "Сегодня мы прощаемся с нашим прекрасным джуном! \nТеперь придется самостоятельно ходить за кофе"

		case "Ладно, пусть работает!":
			msg.Text = "Ура, джун безмерно благодарен и постарается больше не ронять сервер :)"

		case "Послать заказчика":
			msg.Text = "Кому-то прислали кучу правок в выходной день и попросили внести до понедельника? Может пойти на встречу?"
			msg.ReplyMarkup = buttons.CreateDevButtons()

		case "Да ну его, у меня выходной!":
			msg.Text = "Вы правы, можно ответить ему и в понедельник!"

		case "Хорошо, дам ему последний шанс.":
			msg.Text = "Супер, восхищен вашей выдержкой!"

		case "Открыть бизнес":
			msg.Text = "Какой бизнес!? Ахахахах\nИди код писать, дедлайн горит!"
			msg.ReplyMarkup = buttons.CreateBusinessButtons()

		case "У меня отличная идея":
			msg.Text = "Хоть 100, дедлайн горит иди дописывай тесты!!!"

		case "Пойду дописывать тесты...":
			msg.Text = "Отлично! Такими темпами скоро на повышение!"

		default:
			msg.Text = "Прошу прощения! Но я вас не понимаю :(\n С помощью команды /help можно получить список команд"
		}

		_, err := bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}
}
