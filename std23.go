// Quote of the Day Protocol
// Copyright 2015 David Crosby
package main

import (
	"io"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	io.WriteString(conn, "Test test STD 23 has been implemented!\n")
}

func main() {
	ln, err := net.Listen("tcp", ":8017")
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
