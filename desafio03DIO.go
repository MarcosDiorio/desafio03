package main

import (
	"fmt"
	"time"
)

func pin(c1 chan string) {
	for i := 0; ; i++ {
		c1 <- "PING"
	}
}

func pon(c2 chan string) {
	for i := 0; ; i++ {
		c2 <- "PONG"
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c1 chan string = make(chan string)
	var c2 chan string = make(chan string)
	go pin(c1)
	go imprimir(c1)
	go pon(c2)
	go imprimir(c2)

	var entrada string
	fmt.Scanln(&entrada)

}
