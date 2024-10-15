package main

import( 
	"fmt"
)

func swap(a *int, b *int) {
	 var temp int = *a
	 *a = *b
	 *b = temp
}

func sumSlice(slice [] int) *int {

	var sum int =0
	for _,value := range slice {
		sum += value
	}
	return &sum
}

type Person struct {
	name string
	age uint8
}

func newPerson() *Person {
	  return &Person {"murugan", 47}
}

func main() {

	fmt.Println("pointers basics")
	/*
	var a int =10
	var p *int = &a
	fmt.Println("a :" ,a , " *P:" , p)

	var pp **int = &p
	fmt.Println("a :" ,a , " *P:" , *p , "*pp", **pp)
	
	var ppp ***int = &pp
	fmt.Println("a :" ,a , ", *P:" , *p , ", pp:", **pp, ", ppp:", ***ppp)

	***ppp=200
	fmt.Println("a :" ,a , ", *P:" , *p , ", pp:", **pp, ", ppp:", ***ppp)
	*/

	var a int = 100
	var b int = 200
	swap(&a,&b)
	fmt.Println("a :" , a , " b:",b)

	slice := [] int { 2,3,5,6,7 }
	var ptr *int = sumSlice(slice)
	fmt.Println("Sum slice :", *ptr)

	var p Person=Person { "AAA" , uint8(21) }
	fmt.Println(p.name ,":" , p.age)

	 ps := &p
	 ps.name = "BBB"
	 ps.age =20
	fmt.Println(p.name ,":" , p.age)

	var nP *Person = newPerson()
	fmt.Println(nP.name ,":" , nP.age)
	

}