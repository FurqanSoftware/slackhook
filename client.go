package slackhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(body)
	defer resp.Body.Close()
	return nil
}
