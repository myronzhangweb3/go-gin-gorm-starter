package lark

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Lark struct {
	Name   string
	URL    string
	Client *http.Client
}

type MessageContent struct {
	MsgType string     `json:"msg_type"`
	Content PostSchema `json:"content"`
}

type PostSchema struct {
	Post map[string]PostLang `json:"post"`
}

type PostLang struct {
	Title   string        `json:"title"`
	Content [][]TextBlock `json:"content"`
}

type TextBlock struct {
	Tag  string `json:"tag"`
	Text string `json:"text"`
}

// NewLark creates a new Lark webhook client
func NewLark(name, url string) (*Lark, error) {
	if url == "" {
		return nil, errors.New("lark url is empty")
	}
	return &Lark{
		Name: name,
		URL:  url,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}, nil
}

func (l *Lark) send(payload MessageContent) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", l.URL, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := l.Client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("lark: incorrect response status code: %d", resp.StatusCode)
	}

	return nil
}

func (l *Lark) SendMessage(title, details string) error {
	payload := MessageContent{
		MsgType: "post",
		Content: PostSchema{
			Post: map[string]PostLang{
				"zh_cn": {
					Title: l.Name + " Message",
					Content: [][]TextBlock{
						{{Tag: "text", Text: "Type: " + title}},
						{{Tag: "text", Text: "Details: " + details}},
					},
				},
			},
		},
	}
	return l.send(payload)
}

func (l *Lark) SendErrorMessage(title string, err error) error {
	payload := MessageContent{
		MsgType: "post",
		Content: PostSchema{
			Post: map[string]PostLang{
				"zh_cn": {
					Title: l.Name + " Error",
					Content: [][]TextBlock{
						{{Tag: "text", Text: "Error title: " + title}},
						{{Tag: "text", Text: "Error details: " + err.Error()}},
					},
				},
			},
		},
	}
	return l.send(payload)
}
