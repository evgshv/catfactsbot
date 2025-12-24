package botservice

import (
	"net/http"
	"strings"
	"time"
)

type Reciever interface {
	Recieve(client http.Client) (int64, string)
}

type Sender interface {
	Send(id int64, txt string, client http.Client)
}

type Dispatcher interface {
	Reciever
	Sender
}

type BotService struct {
	ChatID int64
	Text   string
}

func New() *BotService {
	return &BotService{}
}

func (p *BotService) Serve(rec Reciever, dis Dispatcher, client http.Client) {

	for {
		var text string

		p.ChatID, p.Text = Dispatcher.Recieve(dis, client)
		if p.Text != "" && strings.HasPrefix(p.Text, "/") {
			command := strings.Split(p.Text, " ")[0]

			switch command {
			case "/start":
				text = "Greetings!"
			case "/getfact":
				_, text = Reciever.Recieve(rec, client)
			}
			if len(text) != 0 {
				Dispatcher.Send(dis, p.ChatID, text, client)
			}
		}
		time.Sleep(1 * time.Second)
	}
}
