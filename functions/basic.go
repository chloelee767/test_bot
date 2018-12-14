package functions

import (
	bot "github.com/chloelee767/test_bot/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var Start = bot.Command{"start", startfn}
var Echo = bot.Command{"echo", echofn}

var startfn = func(tb *bot.TestBot, original *tgbotapi.Message) error {
	tb.SendTextMessage(original, "Welcome")
	return nil
}

var echofn = func(tb *bot.TestBot, original *tgbotapi.Message) error {
	tb.SendTextMessage(original, original.CommandArguments())
	return nil
}
