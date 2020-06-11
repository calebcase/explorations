package main

import (
	"fmt"
	"net"
)

func cannot(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	cannot(err)

	for {
		conn, err := ln.Accept()
		cannot(err)

		fmt.Fprintf(conn, "Silly...\n")
		fmt.Fprintf(conn, "Rabbit")

		// Hold on to this forever...
		//conn.Close()
	}
}
