package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	// загружаем переменные из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	// получаем токен из env
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set")
	}

	// подключаемся к api telegram
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic("Ошибка создания бота:", err)
	}

	// устанавливаем режим отладки
	bot.Debug = true
	log.Printf("Авторизовались как %s", bot.Self.UserName)

	// создаем обработчик обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// запускаем бесконечный цикл обработки обновлений
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		// пропускаем пустые обновления типа пользователь зашел/вышел
		if update.Message == nil {
			continue
		}

		// логируем сообщение в консоль
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// обаботка команды /start
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				// создаем кнопки
				keyboard := tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("🔍 Поиск мема"),
						tgbotapi.NewKeyboardButton("🎲 Рандомный мем"),
					),
				)

				// оправляем приветствие + кнопки
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я бот для поиска мемов!\nВыберите действие:")
				msg.ReplyMarkup = keyboard
				bot.Send(msg)

			case "help":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Это бот, который ищет мемы по ключевому слову.\nВыберите 'Поиск мема' или 'Рандомный мем'.")
				bot.Send(msg)

			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я не знаю такую команду 😢")
				bot.Send(msg)
			}
		} else {
			// обаботка кнопок
			switch update.Message.Text {
				case "🔍 Поиск мема":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Напиши слово, по которому будем искать мем:")
					bot.Send(msg)

				case "🎲 Рандомный мем":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пока просто заглушка бро")
					bot.Send(msg)

				default:
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я не знаю такую команду 😢")
					bot.Send(msg)
			}
		}
	}
}
