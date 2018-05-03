package bot

import (
	"gopkg.in/tucnak/telebot.v2"
    "time"
    "fmt"
)

func Log2Me(message string) {
	tbot.Send(&me, message)
	Log(message)
}

func Log(message string) {
	date := time.Now()
	fmt.Println(date.Format(time.RFC3339), message)
}

func Log2Sender(sender *telebot.User, message string) {
	tbot.Send(sender, message)
	Log(message)
}
