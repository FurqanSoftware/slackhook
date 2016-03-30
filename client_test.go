package slackhook

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNew(t *testing.T) {
	url := "http://example.com/fake-slack-webhook"
	c := New(url)
	if c.url != url {
		t.Fatalf("expected c.url == %q, got %q", url, c.url)
	}
}

func newTrapServer(bodyVar *string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		*bodyVar = string(b)
	}))
}

func TestSend(t *testing.T) {
	body := ""
	srv := newTrapServer(&body)

	cases := []struct {
		msg  Message
		body string
	}{
		{
			msg:  Message{},
			body: `{"text":"","username":"","icon_emoji":"","channel":"","attachments":null}`,
		},
		{
			msg: Message{
				Text: "This is a line of text.\nAnd this is another one.",
			},
			body: `{"text":"This is a line of text.\nAnd this is another one.","username":"","icon_emoji":"","channel":"","attachments":null}`,
		},
		{
			msg: Message{
				Username:  "ghost-bot",
				IconEmoji: ":ghost:",
				Text:      "BOO!",
			},
			body: `{"text":"BOO!","username":"ghost-bot","icon_emoji":":ghost:","channel":"","attachments":null}`,
		},
		{
			msg: Message{
				Text:      "Testing",
				Username:  "ghost",
				IconEmoji: ":ghost:",
				Channel:   "#random",
				Attachments: []Attachment{
					{
						Pretext: "Pretext",
						Color:   "danger",

						AuthorName: "GhostDaddy",
						AuthorLink: "https://www.ghostdad.com",

						Title:     "The Real Ghost",
						TitleLink: "https://www.ghostexample.com",

						Text: "I AM A REAL SLACK GHOST!",
						Fields: []Field{
							{
								Title: "Daddy of all Ghost",
								Value: "The Ghost Daddy",
								Short: true,
							},
						},
					},
				},
			},
			body: `{"text":"Testing","username":"ghost","icon_emoji":":ghost:","channel":"#random","attachments":[{"pretext":"Pretext","color":"danger","author_name":"GhostDaddy","author_link":"https://www.ghostdad.com","author_icon":"","title":"The Real Ghost","title_link":"https://www.ghostexample.com","text":"I AM A REAL SLACK GHOST!","fields":[{"title":"Daddy of all Ghost","value":"The Ghost Daddy","short":true}],"image_url":"","thumb_url":""}]}`,
		},
	}
	for i, c := range cases {
		err := New(srv.URL).Send(c.msg)
		if err != nil {
			t.Fatal(err)
		}
		if body != c.body {
			t.Fatalf("#%d: expected body == %q, got %q", i+1, c.body, body)
		}
	}
}
