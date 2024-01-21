package main

import "fmt"

// fmt is a package from go standard library
// fmt is for formating strings and printing them to the console.

// the name of functions starts with captical characters

func main() {
	// this is the entry point of the application, so when we run our program, 'go' will look fot this main function to execute automatically
	// if this main fun is called with something else instead of 'main' then go will not automatically fire this program

	// there will be only one main function in the application, we need not to have this in every file inside a package
	// In all the files we can create all the functions which can be involved inside this main function, but we only have one main function

	fmt.Println("Hello World!")
}