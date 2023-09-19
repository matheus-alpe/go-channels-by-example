package examples

import (
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	server := NewServer()
	server.Start()

	for i := 0; i < 10; i++ {
		go func(i int) {
			server.userch <- fmt.Sprintf("user %d", i)
			// server.addUser(fmt.Sprintf("user %d", i)) // channels instead map to prevent race conditions
		}(i)
	}

}
