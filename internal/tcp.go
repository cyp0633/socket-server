package internal

import (
	"log"
	"net"

	"go.uber.org/zap"
)

func TCPHandler(conn net.Conn) {
	// 添加客户端
	addr := conn.RemoteAddr().(*net.TCPAddr).IP.String()
	client := Client{Addr: addr, Conn: conn, Messages: make(chan Message, 100)}
	Clients = append(Clients, client)
	Logger.Info("New client", zap.String("addr", addr))

	defer conn.Close()
	for {
		var buf = make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Default().Println(err)
			ProcessExit(client) // 通信中途断开，删除客户端
			return
		}
		Logger.Info("Received", zap.String("addr", addr), zap.String("msg", string(buf[:n])))
		var reply string
		switch msg := string(buf[:n]); {
		case Helo.MatchString(msg):
			reply = ProcessHelo(msg)
		case Send.MatchString(msg):
			reply = ProcessSend(msg, client.Addr)
		case Pull.MatchString(msg):
			reply = ProcessPull(client)
		case User.MatchString(msg):
			reply = ProcessUser()
		case Exit.MatchString(msg):
			conn.Write([]byte(ProcessExit(client)))
			conn.Close()
			return
		default:
			Logger.Warn("Unknown command", zap.String("addr", addr), zap.String("msg", msg))
			reply = "ERROR\n"
		}
		_, err = conn.Write([]byte(reply))
		if err != nil {
			log.Default().Println(err)
			return
		}
	}
}
