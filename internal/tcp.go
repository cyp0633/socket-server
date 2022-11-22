package internal

import (
	"log"
	"net"
)

func TCPHandler(conn net.Conn) {
	defer conn.Close()
	for {
		var buf = make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Default().Println(err)
			return
		}
		log.Default().Println(string(buf[:n]))
		switch msg := string(buf[:n]); {
		case msg == "HELO":
		}
	}

}
