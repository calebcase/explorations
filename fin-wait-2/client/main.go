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

	fmt.Fprintf(conn, "Hello!\n")

	data, err := bufio.NewReader(conn).ReadString('\n')
	cannot(err)

	fmt.Println(data)

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
