package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    uint8
	salary float32
}

// constructor
func newPerson(name string, age uint8, salary float32) *Person {
	return &Person{name: name, age: age, salary: salary}
}

// associate the function with Person
func (p *Person) changeAge(newage uint8) {
	p.age = newage
}

// associate the function with Person
func (p *Person) changeSalary(newsalary float32) {
	p.salary = newsalary
}

// associate the function with Person
func (p *Person) toString() {
	fmt.Println(*p)
}

// inhertance via composition
type Department struct {
	Person
	dept string
}

type Animal interface {
	makeSound()
	isBite() bool
}

type Dog struct{}
type Snake struct{}
type Goat struct{}

func (p Dog) makeSound() {
	fmt.Println("wow wowo woooo")
}
func (p Dog) isBite() bool {
	return true
}
func (p Snake) makeSound() {
	fmt.Println("vozzz vozzz zzz")
}
func (p Snake) isBite() bool {
	return true
}

func (p Goat) makeSound() {
	fmt.Println("meyaww meyaww www")
}
func (p Goat) isBite() bool {
	return false
}

func animalProps(animal Animal) {
	animal.makeSound()
	fmt.Println("Bitting", animal.isBite())
}

/*
func main() {


		var p = Person { "AAA" , uint8(21) , 1000.20 }
		fmt.Println(p)

		 ip := Person { "BBB" , uint8(22) , 2000.20 }
		 fmt.Println(ip)

		 var ptr *Person = &Person { "CCC" , uint8(25) , 5000.20 }
		 fmt.Println(*ptr)

		 iptr := &Person { "ddd" , uint8(35) , 7000.20 }
		 fmt.Println(*iptr)



		ptr  := newPerson("murugan",uint8(48),float32(5000) )
		fmt.Println(*ptr)

		ptr.changeAge(47)
		ptr.changeSalary(7500)

		ptr.toString()


			// inhertance via composition
		// though The toString function associated with Person, it can be called via deptment object.

		var d1 =  Department { Person{"Murugan" , uint8(48), float32(5644.0) },"computer"}
		d1.toString()

		// polymorphisum

		d1 :=Dog{}
		animalProps(d1)
		animalProps(Snake{})
		animalProps(Goat{})


}*/
