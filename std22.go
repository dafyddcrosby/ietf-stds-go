// Character Generator Protocol
// Copyright 2015 David Crosby
package main

import(
	"net"
	"io"
//	"fmt"
)

func handleConnection(conn net.Conn) {
	// According to the RFC any sort of data can be sent, but it's
	// recommended that a recognizable pattern is used "in tha data"
	defer conn.Close()
	var line, col int
	for {
		for col = 0; col < 71; col++ {
			io.WriteString(conn, string('!' + ((col + line) % 94)))
		}
		io.WriteString(conn, "\n")
		line++
		if line == 94 {
			line = 0
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8019")
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
