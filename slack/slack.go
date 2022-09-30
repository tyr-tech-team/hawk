// Package slack provides slack ﳑ
package slack

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// SendMessage - 傳送訊息
func SendMessage(ctx context.Context, c Config, msg string, attachment []Attachment) error {
	// new a Resty client
	client := resty.New()

	// new request
	resp := client.R()

	// set header
	setHeader(resp, c.BotToken)

	req := &Message{
		Channel:     c.Channel,
		Text:        "",
		Attachments: attachment,
		LinkNames:   true,
	}

	// set post body
	resp.EnableTrace().SetBody(req)

	// set response struct
	response := new(MessageResponse)
	resp.SetResult(&response)

	// post
	_, err := resp.Post(fmt.Sprintf("%s%s", SlackHost, SlackSendMessage))
	if err != nil {
		return err
	}

	return nil
}

// setHeader -
func setHeader(r *resty.Request, token string) {
	r.SetHeader("Authorization", "Bearer "+token)
	r.SetHeader("Content-Type", "application/json; charset=utf-8")
}
