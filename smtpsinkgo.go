package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":25")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleSMTPConnection(conn)
	}
}

func handleSMTPConnection(conn net.Conn) {
	defer conn.Close()

	_, err := conn.Write([]byte("250 localhost ESMTP service ready\r\n"))
	if err != nil {
		log.Println(err)
		return
	} else {
		sClientMsg, err := readConn(conn)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(sClientMsg)
	}
}

func readConn(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)

	yNextLine, bHasMoreLines, err := reader.ReadLine()
	if err != nil {
		return "", err
	}
	sMessage := string(yNextLine)

	for bHasMoreLines {
		yNextLine, bHasMoreLines, err = reader.ReadLine()
		if err != nil {
			break
		}

		sMessage += string(yNextLine)
	}

	return sMessage, err
}
