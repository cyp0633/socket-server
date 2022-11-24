package main

import (
	"log"
	"net"
	"socket-server/internal"

	"go.uber.org/zap"
)

func main() {
	ln, err := net.Listen("tcp", ":65432")
	internal.Logger.Info("Listening on :65432")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			internal.Logger.Fatal("Accept error", zap.Error(err))
		}
		go internal.TCPHandler(conn)
	}
}
