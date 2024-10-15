package main

import (
	"fmt"
	"time"
)

func infinitefor(in <-chan string) {

	for {
		select {
		case msg := <-in:
			fmt.Printf("go running finished:%s", msg)
		default:
			fmt.Println("go runing")
		}
	}
}

func main() {

	readOnlyChannel := make(chan string)

	go infinitefor(readOnlyChannel)

	time.Sleep(3 * time.Second)
	readOnlyChannel <- "done"

	fmt.Println("main go finished")
}
