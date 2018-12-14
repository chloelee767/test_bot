package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TestBot struct {
	*tgbotapi.BotAPI
	CommandsMap map[string]CommandFn
}

type CommandFn func(*TestBot, *tgbotapi.Message) error

type Command struct {
	Str string
	Fn  CommandFn
}

type botConfig struct {
	Token string
}

//New cretes new test bot
func New() (TestBot, error) {

	tgbot, err := tgbotapi.NewBotAPI(Token)

	if err != nil {
		log.Fatal(err)
		return TestBot{nil, nil}, err
	}

	bot := TestBot{tgbot, make(map[string]CommandFn)}
	return bot, nil
}

func (tb *TestBot) AddCommand(cmd Command) {
	tb.CommandsMap[cmd.Str] = cmd.Fn
}

func (tb *TestBot) Start(updateFreq int) {
	cfg := tgbotapi.NewUpdate(0)
	cfg.Timeout = updateFreq

	channel, err := tb.GetUpdatesChan(cfg)

	if err != nil {
		log.Fatal(err)
		return
	}

	for update := range channel {

		log.Printf("\nMessage from %s: %s\n\n", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			fn, present := tb.CommandsMap[update.Message.Command()]
			if present {
				err := fn(tb, update.Message)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}

func (tb *TestBot) SendTextMessage(original *tgbotapi.Message, msg string) {
	tb.Send(tgbotapi.NewMessage(original.Chat.ID, msg))
}
