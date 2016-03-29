package slackhook

import "testing"

func TestNew(t *testing.T) {
	url := "http://example.com/fake-slack-webhook"
	c := New(url)
	if c.url != url {
		t.Fatalf("#1: Expected URL: %s, got %s", url, c.url)
	}
}
