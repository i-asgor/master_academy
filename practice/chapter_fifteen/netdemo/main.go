package main

import (
	"fmt"
	"net"
)

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("getting local IP Address..")
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println(err.Error())
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPAddr:
				if v.String() != "0.0.0.0" {
					fmt.Println(v.String())
				}
			}
		}
	}
}
