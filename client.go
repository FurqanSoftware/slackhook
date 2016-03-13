package slackhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	url string
}

func New(url string) *Client {
	return &Client{
		url: url,
	}
}

func (c *Client) Send(msg Message) error {
	b, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
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
