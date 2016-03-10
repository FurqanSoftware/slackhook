package slackhook

type Message struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	Channel   string `json:"channel"`
}
