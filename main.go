package main

import (
	"github.com/chloelee767/test_bot/bot"
	"github.com/chloelee767/test_bot/functions"
)

func main() {
	tb, _ := bot.New()
	tb.Debug = true
	tb.AddCommand(functions.Start)
	tb.AddCommand(functions.Echo)
	tb.Start(60)
}
