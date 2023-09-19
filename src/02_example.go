package examples

import (
	"fmt"
)

func ExecuteExample02() {
	fmt.Println("example 02")
	msgch := make(chan string, 1)

	sendMessage(msgch)
	receiveMessage(msgch)
}

func sendMessage(msgch chan<- string) {
	msgch <- "hello"
}

func receiveMessage(msgch <-chan string)  {
	msg := <-msgch
	fmt.Println(msg)
}
