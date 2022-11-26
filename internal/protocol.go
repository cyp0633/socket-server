package internal

import (
	"fmt"
	"regexp"

	"go.uber.org/zap"
)

// HELO 信息格式，如：HELO
var Helo = regexp.MustCompile(`^HELO`)

// ProcessHelo 返回在线客户端列表
func ProcessHelo(msg string) (clients string) {
	clients = "CLIENTS "
	for _, client := range Clients {
		clients += client.Addr + " "
	}
	clients += "\n"
	Logger.Info("Helo", zap.String("clients", clients))
	return
}

// SEND 信息格式，如：SEND 127.0.0.1 MSG This is Message
var Send = regexp.MustCompile(`^SEND ((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4} MSG .+`)

// 提取 IP 使用
var sendIP = regexp.MustCompile(`^SEND .+ MSG`)

// 提取消息使用
var sendMessage = regexp.MustCompile(`MSG .+`)

// ProcessSend 发送消息
func ProcessSend(msg string, from string) (reply string) {
	temp := sendIP.FindString(msg)
	addr := temp[5 : len(temp)-4]
	content := sendMessage.FindString(msg)[4:]
	message := Message{From: from, To: addr, Content: content}
	for _, client := range Clients {
		if client.Addr == addr {
			client.Messages <- message
			reply = "OK\n"
			Logger.Info("Send", zap.String("from", from), zap.String("to", addr), zap.String("content", content))
			return
		}
	}
	reply = "ERROR\n"
	Logger.Error("Can't find send target", zap.String("from", from), zap.String("to", addr), zap.String("content", content))
	return
}

// PULL 信息格式，如：PULL
var Pull = regexp.MustCompile(`^PULL`)

func ProcessPull(client Client) (reply string) {
	reply = "LEN " + fmt.Sprint(len(client.Messages)) + "\n"
	for len(client.Messages) > 0 {
		msg := <-client.Messages
		reply += "FROM " + msg.From + " CONTENT " + msg.Content + "\n"
	}
	reply += "END\n"
	Logger.Info("Pull", zap.String("client", client.Addr), zap.String("reply", reply))
	return
}

// Exit 信息格式，如：EXIT
var Exit = regexp.MustCompile(`^EXIT`)

// ProcessExit 退出
func ProcessExit(client Client) (reply string) {
	for i, c := range Clients {
		if c.Addr == client.Addr {
			Clients = append(Clients[:i], Clients[i+1:]...)
		}
	}
	reply = "OK\n"
	Logger.Info("Exit", zap.String("client", client.Addr))
	return
}

var User = regexp.MustCompile(`^USER`)

func ProcessUser() (reply string) {
	reply = fmt.Sprintf("%v USERS\n", len(Clients))
	for _, u := range Clients {
		reply += u.Addr + "\n"
	}
	reply += "END\n"
	Logger.Info("User", zap.Any("clients", Clients))
	return
}
