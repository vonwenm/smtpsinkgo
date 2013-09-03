package main

import (
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":21")

	if err != nil {
		// handle error
	}

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
	conn.Close()
}
