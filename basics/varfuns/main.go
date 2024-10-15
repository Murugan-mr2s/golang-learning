package main

import(
	"fmt"
)

// one return type function
func add(a int, b int) int { 
		return (a+b)
}

// two return type function
func divide(a int,  b int) (int,int) {

	quotient := a/b
	reminder := a % b
	return quotient,reminder
}

// no return type function
func printme(name string) {
	fmt.Printf("\nname :%s" , name)
}



func main() {

	var a int
	a = 100

	var c ='A'
	var b =10.45644
	var s = "hello world" 

	fmt.Printf("a :%d , c : %.2f\n",a,b)
	fmt.Printf("s :%s , c : %c\n",s,c)

	name := "murugan"
	age :=47

	fmt.Printf("\nname :%s , age :%d",name,age)

	fmt.Printf("\nAddtion : %d" , add(10,20) )

	q,r :=divide(100,22)
	fmt.Printf("\nDivision : %d ,%d" ,q,r  )
	printme(name)


	towpower := func(x int) int {
		return 2*x
	} 

	fmt.Println("\n 2 power of:", towpower(5))

}