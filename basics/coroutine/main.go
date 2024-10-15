package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func PrintNumbers(wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Println("Coroutine starts")

	for i := 0; i < 10; i++ {
		fmt.Println("N :", i)
		time.Sleep(2 * time.Second)
	}
}

func SendNumbersByChannel(wg *sync.WaitGroup, ch chan int) {

	defer wg.Done()
	fmt.Println("Coroutine starts")
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
	fmt.Println("Coroutine ends")
}

func ChannelSquare(wg *sync.WaitGroup, ch chan int, slice []int) {

	defer wg.Done()
	fmt.Println("Coroutine starts")
	for _, value := range slice {
		ch <- (value * value)
		time.Sleep(1 * time.Second)
	}
	close(ch)
	fmt.Println("Coroutine ends")
}

func ChannelError(wg *sync.WaitGroup, ch chan int, err chan error, slice []int) {

	defer wg.Done()
	fmt.Println("Coroutine starts")
	for _, value := range slice {
		if value == 0 {
			err <- errors.New(" coroutine value =0")
			close(ch)
			return
		}
		ch <- (value * value)
		time.Sleep(1 * time.Second)
	}
	close(ch)
	fmt.Println("Coroutine ends")
}

/*
func main() {

	var wg sync.WaitGroup
	fmt.Println(" Coroutine basics")

	wg.Add(1)
	ch:=make(chan int,10)

	//go PrintNumbers(&wg)

	go SendNumbersByChannel(&wg,ch)
	for num := range ch {
		fmt.Println(num)
	}


	slice := [] int {1,3,5,7,4,0}
	go ChannelSquare(&wg,ch,slice)
	for num := range ch {
		fmt.Println("Square :",  num)
	}


	//slice := [] int {1,3,5,0,4,7}
	errch := make(chan error,2)

	go ChannelError(&wg,ch,errch,slice)
	for num := range ch {
		fmt.Println("Square :",  num)
	}
	wg.Wait()
	err  := <- errch
	if err !=nil {
		fmt.Println("Coroutine finished with error")
	} else {
		fmt.Println("Coroutine finished successfully")
	}
	close(errch)
	fmt.Println("main thread closed")
}
*/
