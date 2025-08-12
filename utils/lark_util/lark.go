package lark_util

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Lark struct {
	name string
	url  string
}

func NewLark(name, url string) (*Lark, error) {
	if url == "" {
		return nil, errors.New("lark_util url is empty")
	}
	return &Lark{
		name: name,
		url:  url,
	}, nil
}

func (l *Lark) messageInfo(title string, info string) string {
	return fmt.Sprintf("{\"msg_type\":\"post\",\"content\":{\"post\":{\"zh_cn\":{\"title\":\"%s\",\"content\":%s}}}}", title, info)
}

func (l *Lark) sendMessage(title string, info string) ([]byte, error) {
	msg := l.messageInfo(l.name+" "+title, info)

	req, err := http.NewRequest("POST", l.url, bytes.NewBuffer([]byte(msg)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("larkutils: incorrect response status code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func (l *Lark) escapeString(s string) string {
	var builder strings.Builder

	for _, r := range s {
		switch r {
		case '\n':
			builder.WriteString("\\n")
		case '\t':
			builder.WriteString("\\t")
		case '\\':
			builder.WriteString("\\\\")
		case '"':
			builder.WriteString("\\\"")
		case '\'':
			builder.WriteString("\\'")
		case '\r':
			builder.WriteString("\\r")
		case '\b':
			builder.WriteString("\\b")
		case '\f':
			builder.WriteString("\\f")
		case '\v':
			builder.WriteString("\\v")
		default:
			builder.WriteRune(r)
		}
	}

	return builder.String()
}

func (l *Lark) SendErrorMessage(title string, error error) ([]byte, error) {
	msg := fmt.Sprintf("["+
		"[{\"tag\":\"text\",\"text\":\"Error title: %s\"}],"+
		"[{\"tag\":\"text\",\"text\":\"Error details: %s\"}]"+
		"]", title, l.escapeString(error.Error()))
	return l.sendMessage("Error Message", msg)
}

func (l *Lark) SendMessage(title string, details string) ([]byte, error) {
	msg := fmt.Sprintf("["+
		"[{\"tag\":\"text\",\"text\":\"Type: %s\"}],"+
		"[{\"tag\":\"text\",\"text\":\"Details: %s\"}]"+
		"]", title, l.escapeString(details))
	return l.sendMessage("Message", msg)
}
