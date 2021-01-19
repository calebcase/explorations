package main

import (
	"fmt"
	"net"
	"os"
	"time"
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

		if os.Getenv("SERVER_FIX_CLOSE_CONN") == "1" {
			// Wait 10 seconds then close the connection.
			go func() {
				time.Sleep(10 * time.Second)
				conn.Close()
			}()
		}
	}
}
