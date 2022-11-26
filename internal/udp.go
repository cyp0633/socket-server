package internal

import (
	"net"
	"regexp"

	"go.uber.org/zap"
)

var udpRegex = regexp.MustCompile(`^PROBE`)

// UDPHandler 处理广播的消息
func UDPHandler(ln *net.UDPConn) {
	data := make([]byte, 1024)
	for {
		n, addr, err := ln.ReadFromUDP(data)
		if err != nil {
			Logger.Fatal("Read error", zap.Error(err))
		}
		if udpRegex.MatchString(string(data[:n])) {
			addr.Port = 65433
			Logger.Info("Client probing", zap.String("addr", addr.String()))
			ln.WriteToUDP([]byte("HERE"), addr)
		}
	}
}
