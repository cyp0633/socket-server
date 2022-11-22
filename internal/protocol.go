package internal

import "regexp"

var Helo = regexp.MustCompile(`^HELO .+`)

// ProcessHelo 返回在线客户端列表
func ProcessHelo(msg string) (clients string) {
	clients = "CLIENTS "
	for _, client := range Clients {
		clients += client.Addr + " "
	}
	return
}
