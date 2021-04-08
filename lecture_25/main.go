package main

import (
	"fmt"
	"time"
)

func main() {

	go cooking("rice ")
	cooking("curry ")

	time.Sleep(time.Second * 10)

}

func cooking(msg string) {
	for i := 0; i <= 10; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}
