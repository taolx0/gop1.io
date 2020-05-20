package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	_, err := io.WriteString(c, time.Now().Format("16:01:03\n"))
	if err != nil {
		log.Fatal(err)
		return
	}
	time.Sleep(1 * time.Second)
}
