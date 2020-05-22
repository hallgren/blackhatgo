package main

import (
	"fmt"
	"io"
	"net"
)

func handle(src net.Conn) {
	defer src.Close()
	dst, err := net.Dial("tcp", "www.internetfirstpage.com:80")
	if err != nil {
		fmt.Println("could not connect", err)
		return
	}
	defer dst.Close()

	go func() {
		_, err := io.Copy(dst, src)
		if err != nil {
			fmt.Println("copy error src -> dst", err)
		}
	}()
	_, err = io.Copy(src, dst)
	if err != nil {
		fmt.Println("copy error dst -> src", err)
	}

}
func main() {
	fmt.Println("proxy")
	socket, err := net.Listen("tcp", ":20081")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := socket.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handle(conn)
	}
}
