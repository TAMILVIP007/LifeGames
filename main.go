package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/amarnathcjd/gogram/telegram"
)

const (
	appID        = 2992000
	appHash      = "235b12e862d71234ea222082052822fd"
	chatusername = "@testeeeeeeeeetchatoo"
	limit        = 1000
)

var client *telegram.Client

func Spambbet() {
	chat, err := client.GetSendablePeer(chatusername)
	if err != nil {
		panic(err)
	}
	for {
		conv1, err := client.NewConversation(chat, false, 30)
		if err != nil {
			fmt.Println(err.Error())
		}
		conv1.SendMessage("/status")
		msg, _ := conv1.GetResponse()
		conv1.Close()
		time.Sleep(2 * time.Second)
		regex := regexp.MustCompile(`Coins: (.+)`)
		ms := regex.FindStringSubmatch(msg.Text())
		if len(ms) < 2 {
			continue
		}
		msgtosend := fmt.Sprintf("/bet %s", strings.ReplaceAll(ms[1][1:], " ", ""))
		i := 1
		for i <= limit {
			client.SendMessage(chat, msgtosend)
			time.Sleep(2 * time.Second)
			i++
		}
// 		conv2, _ := client.NewConversation(chat, false, 30)
// 		conv2.SendMessage("/tier")
// 		msg2, _ := conv2.GetResponse()
// 		conv2.Close()
// 		if msg2.Message.ReplyMarkup == nil {
// 			continue
// 		}
// 		client.MessagesGetBotCallbackAnswer(&telegram.MessagesGetBotCallbackAnswerParams{
// 			Peer:  msg2.Peer,
// 			MsgID: msg2.ID,
// 			Data:  msg2.Message.ReplyMarkup.(*telegram.ReplyInlineMarkup).Rows[0].Buttons[0].(*telegram.KeyboardButtonCallback).Data,
// 		})
// 		time.Sleep(2 * time.Second)
		conv3, _ := client.NewConversation(chat, false, 30)
		conv3.SendMessage(fmt.Sprintf("/deposit %s", strings.ReplaceAll(ms[1], " ", "")))
		msg3, _ := conv3.GetResponse()
		conv3.Close()
		time.Sleep(2 * time.Second)
		regex2 := regexp.MustCompile(`deposit (.+) coins.`)
		ms2 := regex2.FindStringSubmatch(msg3.Text())
		if len(ms2) < 2 {
			continue
		}
		client.SendMessage(chat, fmt.Sprintf("/deposit %s", ms2[1]))
		time.Sleep(2 * time.Second)
	}

}

func main() {
	ok, _ := os.Getwd()
	client, _ = telegram.NewClient(telegram.ClientConfig{
		AppID:    appID,
		AppHash:  appHash,
		LogLevel: telegram.LogInfo,
		Session:  ok + "/session.session",
	})
	if err := client.Start(); err != nil {
		panic(err)
	}
	client.AddMessageHandler("/sfucku", func(message *telegram.NewMessage) error {
		message.Reply("Hello, I am a bot!")
		return nil
	})

	go Spambbet()
	client.Idle()
}
