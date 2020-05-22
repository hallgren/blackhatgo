package main

import (
	"fmt"
	"io"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()
	fmt.Println("echo")
	io.Copy(conn, conn)
}

func main() {
	socket, err := net.Listen("tcp", ":20080")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := socket.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go echo(conn)
	}
}
