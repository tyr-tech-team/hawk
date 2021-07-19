package slack

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
)

const (
	// Slack API Domain
	Slack_Host string = "https://slack.com/api"

	// Slace SendMessage API
	Slack_SendMessage string = "/chat.postMessage"
)

// Message - 傳送訊息
type Message struct {
	// 頻道ID
	Channel string `json:"channel"`
	// 訊息內容
	Text string `json:"text"`
	// 附件數組
	Attachments []Attachment `json:"attachments"`
	// 頭貼
	IconURL string `json:"icon_url"`
	// 機器人名稱
	UserName string `json:"username"`
	// 回覆訊息
	Thread_ts string `json:"thread_ts"`
	// 鏈結(可以tag人或頻道)
	LinkNames bool `json:"link_names"`
}

type Attachment struct {
	MrakdwnIn []string `json:"mrkdwn_in"`
	// 左側的邊框顏色(16進制)
	Color string `json:"color"`
	// 出現在首段的文字
	Pretext string `json:"pretext"`
	// 作者
	AuthorName string `json:"author_name"`
	// 作者連結
	AuthorLink string `json:"author_link"`
	// 作者頭貼
	AuthorIcon string `json:"author_icon"`
	// 標題
	Title string `json:"title"`
	// 標題連結
	TitleLink string `json:"title_link"`
	// 文字介紹
	Text string `json:"text"`
	// 文字區塊
	Fields []TextBlock `json:"fields"`
	// 右側圖像縮圖
	ThumbURL string `json:"thumb_url"`
	// 最底文字
	Footer string `json:"footer"`
	// Footer旁的小圖示
	FooterIcon string `jsno:"footer_icon"`
	// 時間戳
	Ts string `json:"ts"`
}

type TextBlock struct {
	// 標題
	Title string `json:"title"`
	// 內容
	Value string `json:"value"`
	// 並排顯示 (兩個一排)
	Short bool `json:"short"`
}

type MessageResponse struct {
	// 傳送是否成功
	Ok bool `json:"ok"`
	// 傳送的頻道ID
	Channel string `json:"channel"`
	// 該條訊息的回覆ID
	Ts string `json:"ts"`
	// 錯誤訊息
	Error string `json:"error"`
	// 警示
	Warning string `json:"warning"`
}

// SlackConfig -
type SlackConfig struct {
	// 頻道ID
	Channel string `yaml:"channel"`
	// Bot令牌
	BotToken string `yaml:"botToken"`
	// User令牌
	UserToken string `yaml:"userToken"`
}

// SendMessage - 傳送訊息
func SendMessage(ctx context.Context, c SlackConfig, msg string, attachment []Attachment) error {
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
	_, err := resp.Post(fmt.Sprintf("%s%s", Slack_Host, Slack_SendMessage))
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
