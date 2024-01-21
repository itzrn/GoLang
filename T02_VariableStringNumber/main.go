package main

import "fmt"

func main() {

	// string
	// this is one way to declare the variable
	var nameOne string = "Hello Aryan"
	// fmt.Println(nameOne)

	// second way to declare the variable
	var nameTwo = "Hello Man" // here variable type get infer automatcally by the compiler
	// fmt.Println(nameTwo)

	// third way to declare the variable, this is the way where the variable get the default value of its provided type
	var nameTree string
	fmt.Println(nameOne, nameTwo, nameTree)

	// canging the value of variable, but we can not change the data type of variable
	nameOne = "Ishika"
	nameTree = "Ishika Saha" // this is the string given first time to nameThree varaible

	fmt.Println(nameOne, nameTwo, nameTree)

	// another way to declare the variable
	nameFour := "Ishika Aryan" // this can't be used outside the function
	fmt.Println(nameFour)

	// Integer
	var ageOne int = 20
	var ageTwo = 30             // this time go is going to infer the data type
	ageThree := ageOne + ageTwo // this is without using the keyword var

	fmt.Println(ageOne, ageTwo, ageThree)

	// variation of Integer type
	var numOne int8 = 25
	var numTwo int8 = -25
	var numThree uint = 56 // this carry only +ve number and the variation of uint is also available, here in positive we can go more greater than int type as here we have only +ve value
	fmt.Println(numOne, numTwo, numThree)

	// this contains the range of each variation of the data type -> https://pkg.go.dev/builtin

	// same way float data type can also be explored
	e := 2.7 // here the default type is float64
	fmt.Println(e)
}
