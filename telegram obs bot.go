package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"tob/lib/config"
	"tob/plugin/scenes"

	"gopkg.in/telebot.v3"
)

var loadedPlugin []string = []string{}

// https://github.com/tucnak/telebot
// https://github.com/micmonay/keybd_event

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  config.Runtime.BotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	// load plugins
	loadedPlugin = append(loadedPlugin, scenes.Entry(bot))

	// handle basic commands
	bot.Handle("/help", func(c telebot.Context) error {
		rm := &telebot.ReplyMarkup{}
		rm.Inline(
			rm.Row(rm.Text("Control remotely OBS via telegram!")),
			rm.Row(rm.URL("GitHub", "https://github.com/gekigek99/telegram-obs-bot")),
		)
		c.Send("use /menu see commands", rm)
		return nil
	})

	bot.Handle("/menu", func(c telebot.Context) error {
		c.Send(strings.Join(loadedPlugin, "\n"))
		return nil
	})

	fmt.Println("running...")
	bot.Start()
}
