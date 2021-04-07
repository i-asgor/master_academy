package main

import (
	"fmt"
	"time"
)

func main() {
	boring("boring job ")
}

func boring(msg string) {
	for i := 0; i <= 10; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}
