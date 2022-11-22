package internal

import "net"

type Client struct {
	Addr     string
	Conn     net.Conn
	Messages chan Message
}

var Clients = []Client{}

func AddClient(conn net.Conn) {
	addr := conn.RemoteAddr().(*net.TCPAddr).IP.String()
	Clients = append(Clients, Client{Addr: addr, Conn: conn, Messages: make(chan Message)})
}
