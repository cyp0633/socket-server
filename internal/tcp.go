package internal

import (
	"log"
	"net"
)

func TCPHandler(conn net.Conn) {
	// 添加客户端
	addr := conn.RemoteAddr().(*net.TCPAddr).IP.String()
	client := Client{Addr: addr, Conn: conn, Messages: make(chan Message)}
	Clients = append(Clients, client)

	defer conn.Close()
	for {
		var buf = make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Default().Println(err)
			return
		}
		log.Default().Println(string(buf[:n]))
		var reply string
		switch msg := string(buf[:n]); {
		case Helo.MatchString(msg):
			reply = ProcessHelo(msg)
		case Send.MatchString(msg):
			reply = ProcessSend(msg)
		}
		_, err = conn.Write([]byte(reply))
		if err != nil {
			log.Default().Println(err)
			return
		}
	}
}
