package examples

import (
	"fmt"
	"time"
)

func ExecuteExample01() {
	fmt.Println("example 01")
	userch := make(chan string, 1)

	userch <- "bob"

	fmt.Println(<-userch)

	go func() {
		time.Sleep(2 * time.Second)
		userch <- "matt" 
	}()

	fmt.Println("a")
	user := <-userch // block
	fmt.Println("b")

	fmt.Println(user)
}
