package internal

import "net"

type Client struct {
	Addr     string
	Conn     net.Conn
	Messages chan Message
}

var Clients = []Client{}
