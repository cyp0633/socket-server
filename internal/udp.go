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
		Logger.Info("Received", zap.String("addr", addr.String()), zap.String("msg", string(data[:n])))
		if udpRegex.MatchString(string(data[:n])) {
			ln.WriteToUDP([]byte("HERE"), addr)
		}
	}
}
