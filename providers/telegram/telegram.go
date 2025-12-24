package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type TelegramProvider struct {
	ProviderID string
	host       string
	botAPI     string
	offset     int
}

type Update struct {
	ID      int      `json:"update_id"`
	Message *Message `json:"message"`
}

type Message struct {
	Chat struct {
		ID int64 `json:"id"`
	} `json:"chat"`
	Text string `json:"text"`
}

func New(host string, botAPI string, offset int) *TelegramProvider {
	return &TelegramProvider{ProviderID: "telegram", host: host, botAPI: botAPI, offset: offset}
}

func (p *TelegramProvider) updates() ([]Update, error) {
	url := fmt.Sprintf("%sgetUpdates?timeout=60&offset=%d", (p.host + "bot" + p.botAPI + "/"), p.offset)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Ok     bool     `json:"ok"`
		Result []Update `json:"result"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result.Result, nil
}

func (p *TelegramProvider) sendMessage(chatID int64, text string) error {
	url := fmt.Sprintf("%ssendMessage?chat_id=%d&text=%s", (p.host + "bot" + p.botAPI + "/"), chatID, url.QueryEscape(text))
	_, err := http.Get(url)
	return err
}

func (p *TelegramProvider) Recieve() (int64, string) {
	for {
		updates, err := p.updates()
		if err != nil {
			continue
		}
		for _, update := range updates {
			p.offset = update.ID + 1
			return update.Message.Chat.ID, update.Message.Text
		}
	}

}

func (p *TelegramProvider) Send(chatID int64, txt string) {
	p.sendMessage(chatID, txt)
}
