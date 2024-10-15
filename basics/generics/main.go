package main

import (	
			//"os"
			"fmt"	
		)

		// generic function
func Print[T any](input T ) {
		fmt.Println(input)
}

func Back[T any](a T) T {
	 return a
}

func PrintTypes[T,S any] (a T,b S) {
	fmt.Println("T a :", a)
	fmt.Println("T b :", b)
} 

type Collection[T any] struct {
	items []T
}

func (c *Collection[T]) addNew (a T) {
		c.items = append(c.items,a)
}

func (c *Collection[T]) getAll() []T {
		return c.items
}


type Number interface {
	int | int64 | float64 
}

func Sum[T Number] (a,b T) T {
		return (a+b)
}

func main() {

	/*
	os.Stdout.WriteString("Generics Basics")
	Print(499)
	Print("hello world")
	Print("3.142")
	*/

	/*
	// pass general type, return specific type
	var a int =Back(40)  
	b :=Back(40.2)
	var s string =Back("Hello")
	fmt.Println(a,b,s)
	*/

	/*
	PrintTypes(230,100.9)
	PrintTypes("Hello world",100.9)
	*/

	/*
	var C = Collection[int]{}
	C.addNew(10)
	C.addNew(12)
	C.addNew(14)
	fmt.Println(C.getAll())

	S := Collection[string]{}
	S.addNew("AAA")
	S.addNew("FFF")
	S.addNew("DDDD")
	fmt.Println(S.getAll())
	*/

	s:=Sum(10,30)	
	fmt.Println("Sum: ", s)

	s2:=Sum(10.5,30.8)	
	fmt.Println("Sum: ", s2)

}