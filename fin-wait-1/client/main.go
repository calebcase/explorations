package main

import (
	"bufio"
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
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	cannot(err)

	if os.Getenv("CLIENT_FIX_SET_LINGER") == "1" {
		// Discard any unsent data:
		// https://golang.org/pkg/net/#TCPConn.SetLinger
		tconn := conn.(*net.TCPConn)
		tconn.SetLinger(0)
	}

	fmt.Fprintf(conn, "Hello!\n")

	data, err := bufio.NewReader(conn).ReadString('\n')
	cannot(err)

	fmt.Println(data)

	// Write a bunch of data so that our queues get filled up.
	go func() {
		for i := 0; i < 1048576; i++ {
			fmt.Fprintf(conn, ".")
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)

		p, err := os.FindProcess(os.Getpid())
		cannot(err)

		err = p.Signal(os.Interrupt)
		cannot(err)
	}()

	data, err = bufio.NewReader(conn).ReadString('\n')
	cannot(err)

	fmt.Println(data)
}
