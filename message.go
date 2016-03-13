package slackhook

type Message struct {
	Text        string       `json:"text"`
	Username    string       `json:"username"`
	IconEmoji   string       `json:"icon_emoji"`
	Channel     string       `json:"channel"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Pretext string `json:"pretext"`

	Color string `json:"color"`

	AuthorName string `json:"author_name"`
	AuthorLink string `json:"author_link"`
	AuthorIcon string `json:"author_icon"`

	Title     string `json:"title"`
	TitleLink string `json:"title_link"`

	Text string `json:"text"`

	Fields []Field `json:"fields"`

	ImageURL string `json:"image_url"`
	ThumbURL string `json:"thumb_url"`
}

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}
