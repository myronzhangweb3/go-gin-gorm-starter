package lark

import (
	"fmt"
	"testing"
)

func TestSendMessageHello(t *testing.T) {
	// Create a new Lark client with webhook URL
	lark, _ := NewLark("TestMessage", "https://open.larksuite.com/open-apis/bot/v2/hook/3a62d546-e2be-4549-8e7b-bd7f5457253b")

	// Test 1: send a simple English message
	err := lark.SendMessage("test", "Hello World")
	if err != nil {
		t.Fatalf("failed to send normal message: %v", err)
	}

	// Test 2: message with special characters, including Chinese, emoji, quotes, backslash, JSON symbols, URL
	specialText := `Special characters: 中文 ✅, emoji 😊, quotes " ' \, JSON { } [ ], URL: https://example.com`
	err = lark.SendMessage("special", specialText)
	if err != nil {
		t.Fatalf("failed to send special character message: %v", err)
	}

	// Test 3: message with multiple lines and newline characters
	multiLine := "Line one\nLine two: contains newline\nLine three: end."
	err = lark.SendMessage("multiline", multiLine)
	if err != nil {
		t.Fatalf("failed to send multiline message: %v", err)
	}
}

func TestSendMessageError(t *testing.T) {
	// Create a new Lark client with webhook URL for error message test
	lark, _ := NewLark("TestMessageError", "https://open.larksuite.com/open-apis/bot/v2/hook/3a62d546-e2be-4549-8e7b-bd7f5457253b")

	// Test 1: send a simple error message
	err := lark.SendErrorMessage("test error", fmt.Errorf("hello error"))
	if err != nil {
		t.Fatalf("failed to send error message: %v", err)
	}

	// Test 2: send an error message with special characters
	specialErr := fmt.Errorf("Special error: \"quotes\", backslash \\, emoji 🚀, newline\nSecond line")
	err = lark.SendErrorMessage("special error", specialErr)
	if err != nil {
		t.Fatalf("failed to send special error message: %v", err)
	}
}
