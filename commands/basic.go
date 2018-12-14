package commands

import (
	bot "github.com/chloelee767/test_bot/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var Start = func(tb *bot.TestBot, original *tgbotapi.Message) error {
	tb.SendTextMessage(original, "Welcome")
	return nil
}

var Echo = func(tb *bot.TestBot, original *tgbotapi.Message) error {
	tb.SendTextMessage(original, original.CommandArguments())
	return nil
}
