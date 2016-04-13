// Package slackhook provides a client implementation for Slack's
// incoming webhooks.
package slackhook

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// Client represents a Slack client for a single incoming webhook URL.
type Client struct {
	url string
}

// New takes a Slack incoming webhook URL and returns a new Client.
func New(url string) *Client {
	return &Client{
		url: url,
	}
}

// Send takes a Slack message and dispatches it over the webhook URL.
func (c *Client) Send(msg Message) error {
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	resp, err := http.Post(c.url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	_, err = io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
