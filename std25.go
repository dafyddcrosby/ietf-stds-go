// Daytime Protocol
// Copyright 2015 David Crosby

package main

import(
	"net"
	"io"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	io.WriteString(conn, string(time.Now().Format("Monday, January 2, 2006, 15:04:05-MST")))
	io.WriteString(conn, "\n")
}

func main() {
	ln, err := net.Listen("tcp", ":8013")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		go handleConnection(conn)
	}

}
