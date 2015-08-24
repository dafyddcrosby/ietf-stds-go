// Discard Protocol
// Copyright 2015 David Crosby
package main

import(
	"net"
)

func handleConnection(conn net.Conn) {
}

func main() {
	ln, err := net.Listen("tcp", ":8009")
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
