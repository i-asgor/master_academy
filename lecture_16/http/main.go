package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//nl, err := net.Listen("tcp", ":8888") tcp akta confirm protocol jeta amk ensurity ba nischoyota dei je obossoi opor end e message send hobe
	nl, err := net.Listen("tcp", "localhost:8888") // Listen method duti parameter recieved kore orthat duti argument pass korte parbo prothomti network ebong ditioti sei network er access
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//nl.Accept() jodi kono request ase seta amra accept korbo accept method duti data return kore net.conn and error

	for {
		conn, err := nl.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		bs := make([]byte, 1024) //amra jekono normal string ke bits ba bytes e convert korar jonno byte package ta babohar kora hoi. network theke jehetu amra sorasori request nicci request ta by default byte format e thake
		//conn.read() je message amra accept korci seta read korbo
		n, err := conn.Read(bs) // read method byte message recieved kore
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(n)

		reqstr := string(bs)
		fmt.Println(reqstr) //conversion = jokhon duto bhinno bhinno data type er modhe data exchange hoi. r aki dhoroner datar modhe hole seta assertion

		body := `<!DOCTYPE html>
	<html>
	<head>
	<title>Page Title</title>
	</head>
	<body>
	
	<h1>My First Heading</h1>
	<p>My first paragraph.</p>
	
	</body>
	</html>`

		resp := "HTTP/1.1 200 OK\r\n" +
			"Content-Length:%d\r\n" +
			"Content-Type:text/html\r\n" +
			"\r\n%s"
		msg := fmt.Sprintf(resp, len(body), body)
		fmt.Println(msg)
		conn.Write([]byte(msg))
	}
}
