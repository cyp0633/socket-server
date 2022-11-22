package main

import (
	"log"
	"net"
	"socket-server/internal"
)

func main() {
	ln, err := net.Listen("tcp", ":65432")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		go internal.TCPHandler(conn)
	}
}
