package main

import (
	"fmt"
)

func SliceSquare(slice [] int) []int {

	rslice :=[] int{}
	for _,value :=range slice {
		rslice = append(rslice,value*value)
	}
	return rslice
}

func PrintSlice(slice [] int) {
	for index,value :=range slice {
		fmt.Println(index ,":" , value)
	}
}

func findEven(slice [] int) [] int {

	evenSlice := [] int {}
	for _,value := range slice {

			if value%2 == 0  {
				evenSlice =append(evenSlice,value)
			}
	}
	return evenSlice
}

func Sum(slice [] int) int {
	sum :=0
	for _,value := range slice {
		sum += value
	}
	return sum
}

func Average(slice [] int) float64 {
	sum :=0.0
	for _,value := range slice {
		sum = sum + float64(value)
	}
	return sum/ float64(len(slice)) 
}

func splitEvenAndOddSlice(slice [] int) ([]int,[]int) {
	oddSlice := [] int {}
	evenSlice := [] int {}

	for _,value := range slice {
		if value%2 ==0 {
			evenSlice = append(evenSlice,value)
		 } else {
			oddSlice = append(oddSlice,value)
		 }
	}
	return evenSlice,oddSlice
}


func dayOfWeek(slice [] int) [] string {
	days := [] string {}

	for _,value := range slice {

		switch(value) {

		case 0 :
			days = append(days , "Sunday")
		case 1 : 
			days = append(days , "Monday")
		case 2 :
			days = append(days , "Tuesday")
		case 3 :
			days = append(days , "Wednesday")
		case 4 : 
			days = append(days , "Thursday")
		case 5 :	
			days = append(days , "Friday")
		case 7 :
			days = append(days , "Saturday")					
		default :
			days = append(days , "noday")

		}
	}
	return days
}

func main() {

	fmt.Println("arrays condition loop")
	/*	
	var arr [3]int  // declare with fixed dimension
	arr[0]=10
	arr[1]=20
	arr[2]=30

	for index,value := range arr {
		fmt.Println(index , ":" , value)
	} 

	arr2 :=[3]int {2,4,6}  // array by inference with dimension
	for i:=0; i<len(arr2) ; i++ {
		fmt.Println(arr2[i])
	}

	slice := [] int {5,3,2,1,9,4,0}
	slice = append(slice,7,8)
	PrintSlice(slice)
	
	slice2 := slice[2:]
	PrintSlice(slice2)
	
	ss := SliceSquare(slice)
	PrintSlice(ss)
	*/

	slice2 := [] int {5,3,2,1,9,4,0}
	evenSlice := findEven(slice2)
	PrintSlice(evenSlice)

	sum := Sum(slice2)
	ave := Average(slice2)
	fmt.Println("Sum :",sum)
	fmt.Println("Average :",ave)

	evenSlice,oddSlice :=splitEvenAndOddSlice(slice2)

	fmt.Println("Even :", evenSlice)
	fmt.Println("odd :",oddSlice)

	d1 :=dayOfWeek(evenSlice)
	fmt.Println("Even Day :", d1)

	d2 :=dayOfWeek(oddSlice)
	fmt.Println("Odd Day :", d2)
	
}