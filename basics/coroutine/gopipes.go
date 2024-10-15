package main

import "fmt"

func readPipe(data []int) <-chan int {
	out := make(chan int)

	go func() {
		for value, _ := range data {
			out <- value
		}
		close(out)
	}()
	return out
}

func processPipe(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {

		for value := range in {
			out <- value * value
		}
		close(out)
	}()
	return out
}

func main() {

	data := []int{1, 2, 3, 4, 5}

	processChannel := readPipe(data)

	OutputChannel := processPipe(processChannel)

	for output := range OutputChannel {
		fmt.Println(output)
	}

}
