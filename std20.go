// Echo Protocol
// Copyright 2015 David Crosby
package main

import(
	"net"
	"io"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	io.Copy(conn, conn)
}

func main() {
	ln, err := net.Listen("tcp", ":8007")
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
