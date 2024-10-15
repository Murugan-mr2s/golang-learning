package main

import "fmt"

func Squarefunc(a int) int {
	return a * a
}

func InputHfunc(slice []int, Square func(a int) int) {
	for _, v := range slice {
		fmt.Println(Square(v))
	}
}

func OutputHfunc(multi int) func(a int) int {

	return func(a int) int {
		return multi * a
	}
}

func main() {

	vals := []int{1, 2, 3, 4}

	fmt.Println("A function passed as a parameter")
	InputHfunc(vals, Squarefunc)

	fmt.Println("A function returned from method as parameter")
	multiplier := OutputHfunc(10)

	for _, v := range vals {
		fmt.Println(multiplier(v))
	}

}
