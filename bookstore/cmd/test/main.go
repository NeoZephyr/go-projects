package main

import (
	"log"
	"time"
)

func main() {
	println("hello test")
	c := make(chan string)

	select {
	case s := <- c:
		log.Println("channel value:", s)
	case <-time.After(10 * time.Second):
		log.Println("time out")
	}
}
