// Active Users
// This service is a bad idea, but whatevs. Don't run it!
// Copyright 2015 David Crosby

package main

import (
	"io"
	"net"
	"os/exec"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	out, err := exec.Command("who").Output()
	if err != nil {
		conn.Close()
		return
	}
	io.WriteString(conn, string(out))
}

func main() {
	ln, err := net.Listen("tcp", ":8011")
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
