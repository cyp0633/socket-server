package internal

import "net"

type Client struct {
	Addr string
	Conn net.Conn
}

var Clients map[string]net.Conn

func AddClient(conn net.Conn) {
	addr := conn.RemoteAddr().(*net.TCPAddr).IP.String()
	Clients[addr] = conn
}
