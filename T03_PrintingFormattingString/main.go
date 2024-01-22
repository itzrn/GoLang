package main

import "fmt"

func main() {
	age := 21
	name := "Aryan"

	// print
	fmt.Print("Hello, ")   // this won't add new line
	fmt.Print("world! \n") // \n will add new line
	fmt.Print("new Line \n")

	fmt.Println("Hello, ") // this add new line
	fmt.Print("World \n")

	// Println
	fmt.Println("my age is", age, "and my name is", name) // ',' by itself add one space between the words

	// Printf(formatted strings) -> using %v for fomatting specifier
	fmt.Printf("my age is %v and my name is %v\n", age, name)

	// %q puts the string in double quotes, mind it only string literals
	fmt.Printf("my age is %q and my name is %q\n", age, name)

	// %T is used to get the type of the value
	fmt.Printf("age is of type %T\n", age)

	// this is used for the float type value

	fmt.Printf("You scored %f points! \n", 222.55)

	// %0.2f means -> till 2 decimal place value
	fmt.Printf("You scored %0.2f points! \n", 222.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("The saved string is ->", str)

	// to check out more formate specifier -> https://pkg.go.dev/fmt
}
