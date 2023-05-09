package buttons

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// Кнопки со старта
func CreateButtons() tgbotapi.ReplyKeyboardMarkup {
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

// Для любителей кальяна
func CreateHookahButtons() tgbotapi.ReplyKeyboardMarkup {
	buttonRow1 := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("Мне покрепче!"),
	}

	buttonRow2 := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("Стандартно и побыстрее!"),
	}

	buttons := [][]tgbotapi.KeyboardButton{buttonRow1, buttonRow2}

	replyKeyboardMarkup := tgbotapi.NewReplyKeyboard(buttons...)
	replyKeyboardMarkup.OneTimeKeyboard = true

	return replyKeyboardMarkup
}

// Для ненавистников джунов
func CreateJuniorButtons() tgbotapi.ReplyKeyboardMarkup {
	buttonRow1 := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("Да, надо уволить!"),
	}

	buttonRow2 := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("Ладно, пусть работает!"),
	}

	buttons := [][]tgbotapi.KeyboardButton{buttonRow1, buttonRow2}

	replyKeyboardMarkup := tgbotapi.NewReplyKeyboard(buttons...)
	replyKeyboardMarkup.OneTimeKeyboard = true

	return replyKeyboardMarkup
}

// Для тех кого обидел заказчик)
func CreateDevButtons() tgbotapi.ReplyKeyboardMarkup {
	buttonRow1 := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("Да ну его, у меня выходной!"),
	}

	buttonRow2 := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("Хорошо, дам ему последний шанс."),
	}

	buttons := [][]tgbotapi.KeyboardButton{buttonRow1, buttonRow2}

	replyKeyboardMarkup := tgbotapi.NewReplyKeyboard(buttons...)
	replyKeyboardMarkup.OneTimeKeyboard = true

	return replyKeyboardMarkup
}

// Для тех, кто устал писать код
func CreateBusinessButtons() tgbotapi.ReplyKeyboardMarkup {
	buttonRow1 := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("У меня отличная идея"),
	}

	buttonRow2 := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("Пойду дописывать тесты..."),
	}

	buttons := [][]tgbotapi.KeyboardButton{buttonRow1, buttonRow2}

	replyKeyboardMarkup := tgbotapi.NewReplyKeyboard(buttons...)
	replyKeyboardMarkup.OneTimeKeyboard = true

	return replyKeyboardMarkup
}
