package scenes

import (
	"fmt"
	"tob/lib/utils"

	"github.com/micmonay/keybd_event"
	"gopkg.in/telebot.v3"
)

var entryCom string = "/scenes"

// Entry initializes the scenes command control.
//
// Returns the string command to access the menu.
func Entry(bot *telebot.Bot) string {
	sceneTable := &telebot.ReplyMarkup{}
	s1 := sceneTable.Data("1", "1")
	s2 := sceneTable.Data("2", "2")
	s3 := sceneTable.Data("3", "3")
	s4 := sceneTable.Data("4", "4")
	s5 := sceneTable.Data("5", "5")
	sceneTable.Inline(
		sceneTable.Row(s1, s2),
		sceneTable.Row(s3, s4),
		sceneTable.Row(s5),
	)

	bot.Handle(entryCom, func(c telebot.Context) error {
		c.Send("select scene", sceneTable)
		return nil
	})

	bot.Handle(&s1, func(c telebot.Context) error {
		utils.KeyPress(keybd_event.VK_1)
		c.Respond(&telebot.CallbackResponse{Text: "set to 1"})
		return nil
	})

	bot.Handle(&s2, func(c telebot.Context) error {
		utils.KeyPress(keybd_event.VK_2)
		c.Respond(&telebot.CallbackResponse{Text: "set to 2"})
		return nil
	})

	bot.Handle(&s3, func(c telebot.Context) error {
		utils.KeyPress(keybd_event.VK_3)
		c.Respond(&telebot.CallbackResponse{Text: "set to 3"})
		return nil
	})

	bot.Handle(&s4, func(c telebot.Context) error {
		utils.KeyPress(keybd_event.VK_4)
		c.Respond(&telebot.CallbackResponse{Text: "set to 4"})
		return nil
	})

	bot.Handle(&s5, func(c telebot.Context) error {
		utils.KeyPress(keybd_event.VK_5)
		c.Respond(&telebot.CallbackResponse{Text: "set to 5"})
		return nil
	})

	fmt.Println("loaded:", entryCom)

	return entryCom
}
