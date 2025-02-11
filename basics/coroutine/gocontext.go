package main

import (
	"context"
)

func Generator(ctx context.Context, genvalue func() int) <-chan int {
	stream := make(chan int)

	go func() {

		defer close(stream)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				stream <- genvalue()
			}
		}
	}()
	return stream
}

/*
func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	randfunc := func() int {
		return rand.Intn(1000)
	}

	randStream := Generator(ctx, randfunc)

	for v := range randStream {
		fmt.Println(v)
	}

}
*/
/*
func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	randfunc := func() int {
		return rand.Intn(1000)
	}
	randStream := Generator(ctx, randfunc)
	for v := range randStream {
		fmt.Println(v)
		if v%100 == 0 {
			break // break for loop and call defer cancel
		}
	}

}
*/
