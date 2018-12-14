package main

import (
	"github.com/chloelee767/test_bot/bot"
	"github.com/chloelee767/test_bot/commands"
)

func main() {
	tb, _ := bot.New()
	tb.Debug = true
	tb.AddCommand("start", commands.Start)
	tb.AddCommand("echo", commands.Echo)
	tb.Start(60)
}
