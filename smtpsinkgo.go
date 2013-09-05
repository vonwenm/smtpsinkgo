package main

import (
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":25")

	if err != nil {
		// handle error
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			// handle error
			continue
		}

		go handleSMTPConnection(conn)
	}
}

func handleSMTPConnection(conn net.Conn) {
	defer conn.Close()

	_, err := conn.Write([]byte("250 localhost ESMTP service ready\r\n"))
	if err != nil {
		// handle error
		return
	}
}
