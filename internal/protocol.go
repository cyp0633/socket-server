package internal

import (
	"regexp"
	"strings"
)

// HELO 信息格式，如：HELO
var Helo = regexp.MustCompile(`^HELO$`)

// ProcessHelo 返回在线客户端列表
func ProcessHelo(msg string) (clients string) {
	clients = "CLIENTS "
	for _, client := range Clients {
		clients += client.Addr + " "
	}
	return
}

// SEND 信息格式，如：SEND 127.0.0.1 MSG This is Message
var Send = regexp.MustCompile(`/^SEND [0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3} MSG .+`)

// 提取消息使用
var sendMessage = regexp.MustCompile(`MSG .+`)

// ProcessSend 发送消息
func ProcessSend(msg string, from string) (reply string) {
	temp := strings.Split(msg, "")
	addr := temp[1]
	content := sendMessage.FindString(msg)[4:]
	message := Message{From: from, To: addr, Content: content}
	for _, client := range Clients {
		if client.Addr == addr {
			client.Messages <- message
			reply = "OK"
		}
	}
	return
}
