package lark_util

import (
	"fmt"
	"testing"
)

func TestSendMessageHello(t *testing.T) {
	lark, _ := NewLark("TestMessage", "https://open.larksuite.com/open-apis/bot/v2/hook/3a62d546-e2be-4549-8e7b-bd7f5457253b")
	sendMessage, err := lark.SendMessage("test", "Hello World")
	if err != nil {
		panic(err)
	}
	println(string(sendMessage))
}

func TestSendMessageError(t *testing.T) {
	lark, _ := NewLark("TestMessageError", "https://open.larksuite.com/open-apis/bot/v2/hook/3a62d546-e2be-4549-8e7b-bd7f5457253b")
	sendMessage, err := lark.SendErrorMessage("test", fmt.Errorf("hello error"))
	if err != nil {
		panic(err)
	}
	println(string(sendMessage))
}
