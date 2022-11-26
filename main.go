package main

import (
	"net"
	"socket-server/internal"
	"sync"

	"go.uber.org/zap"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go listenTCP(&wg)
	go listenUDP(&wg)
	wg.Wait() // 等待下面两个函数协程执行完毕
}

func listenTCP(wg *sync.WaitGroup) {
	defer wg.Done()
	ln, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 65432,
	})
	if err != nil {
		internal.Logger.Fatal("Listen error", zap.Error(err))
		panic(err)
	}
	internal.Logger.Info("Listening messaging service on 0.0.0.0:65432/tcp")

	for {
		conn, err := ln.Accept()
		if err != nil {
			internal.Logger.Fatal("Accept error", zap.Error(err))
		}
		go internal.TCPHandler(conn)
	}
}

func listenUDP(wg *sync.WaitGroup) {
	defer wg.Done()
	ln, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 65432,
	})
	if err != nil {
		internal.Logger.Fatal("Listen error", zap.Error(err))
	}
	internal.Logger.Info("Listening probe on 0.0.0.0:65432/udp")
	internal.UDPHandler(ln)
}
