package main

import (
	"fmt"
	"net"
)

func main() {
	var conn net.Conn
	var err error
	conn, err = net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println(err.Error())
	}

	//way-1
	//bs := []byte("Hello")
	//conn.Write(bs)

	//Way-2
	conn.Write([]byte("Hello\n"))
	conn.Close()

}
