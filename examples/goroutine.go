package main

import (
	"fmt"
	"time"
)

// document: https://golangbot.com/channels/

// <- chan - read from chan
// chan <- "xxx" - write to chan
func main() {
	fmt.Println("Start...")
	// create channel for string support
	c := make(chan string)

	// using go routine to put data to channel
	go func(c chan string) {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			c <- "Hello World!"
		}
		close(c)
	}(c)

	// forever loop for received message from channel
	for {
		msg, ok := <-c
		if ok {
			fmt.Println(msg)
		} else {
			// break the loop after no data from channel
			break
		}
	}

	fmt.Println("Done!")
}
