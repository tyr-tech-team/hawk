package slack

const (
	// SlackHost - API Domain
	SlackHost string = "https://slack.com/api"

	// SlackSendMessage - API
	SlackSendMessage string = "/chat.postMessage"
)

// Message - 傳送訊息
type Message struct {
	// Channel - 頻道ID
	Channel string `json:"channel"`

	// Text - 訊息內容
	Text string `json:"text"`

	// Attachments - 附件數組
	Attachments []Attachment `json:"attachments"`

	// IconURL - 頭貼
	IconURL string `json:"icon_url"`

	// UserName - 機器人名稱
	UserName string `json:"username"`

	// ThreadTS - 回覆訊息
	ThreadTS string `json:"thread_ts"`

	// LinkNames - 鏈結(可以tag人或頻道)
	LinkNames bool `json:"link_names"`
}

// Attachment -
type Attachment struct {
	// MarkdownIn -
	MarkdownIn []string `json:"markdown_in"`

	// Color - 左側的邊框顏色(16進制)
	Color string `json:"color"`

	// Pretext - 出現在首段的文字
	Pretext string `json:"pretext"`

	// AuthorName - 作者
	AuthorName string `json:"author_name"`

	// AuthorLink - 作者連結
	AuthorLink string `json:"author_link"`

	// AuthorIcon - 作者頭貼
	AuthorIcon string `json:"author_icon"`

	// Title - 標題
	Title string `json:"title"`

	// TitleLink - 標題連結
	TitleLink string `json:"title_link"`

	// Text - 文字介紹
	Text string `json:"text"`

	// Fields - 文字區塊
	Fields []TextBlock `json:"fields"`

	// ThumbURL - 右側圖像縮圖
	ThumbURL string `json:"thumb_url"`

	// Footer - 最底文字
	Footer string `json:"footer"`

	// FooterIcon - Footer旁的小圖示
	FooterIcon string `json:"footer_icon"`

	// Ts - 時間戳
	Ts string `json:"ts"`
}

// TextBlock -
type TextBlock struct {
	// Title - 標題
	Title string `json:"title"`

	// Value - 內容
	Value string `json:"value"`

	// Short - 並排顯示 (兩個一排)
	Short bool `json:"short"`
}

// MessageResponse -
type MessageResponse struct {
	// Ok - 傳送是否成功
	Ok bool `json:"ok"`

	// Channel - 傳送的頻道ID
	Channel string `json:"channel"`

	// Ts - 該條訊息的回覆ID
	Ts string `json:"ts"`

	// Error - 錯誤訊息
	Error string `json:"error"`

	// Warning - 警示
	Warning string `json:"warning"`
}

// Config -
type Config struct {
	// Channel - 頻道ID
	Channel string `yaml:"channel"`

	// BotToken - Bot令牌
	BotToken string `yaml:"botToken"`

	// UserToken - User令牌
	UserToken string `yaml:"userToken"`

	// Open - 開關 0->關閉,1->啟用
	Open int `yaml:"open"`
}
