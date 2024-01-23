// for the std library in golang -> https://pkg.go.dev/std#stdlib
package main

import (
	"fmt"
	"strings"
	"sort"
)

func main()  {
	greeting := "hello world!"

	// using the librabry strings
	fmt.Println(strings.Contains(greeting, "hello"))

	// this won't change the actual variable value until and unless the new value won't get assign to it
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))

	// value won't got change after using the Replace function of strings
	fmt.Println(greeting)

	// to make the string in capital letters, nit changing the value of the variable
	fmt.Println(strings.ToUpper(greeting))

	// this is to find the index of any character or word in string, it gives the first occurance of the giving character or words
	fmt.Println(strings.Index(greeting, "ll"))

	// this is going to split the string where ever the particular character found
	fmt.Println(strings.Split(greeting, " "))

	// creating slices
	ages := []int{8, 6, 7, 5, 4, 2, 3, 1}
	sort.Ints(ages) // this is the way we sort int slices, this change the actual variable value in sorted way
	fmt.Println(ages)

	// this is used to find the idnex of the given integer and if the value won't exist then it returns the length of the slices which is not a valid index
	index := sort.SearchInts(ages, 6)
	fmt.Println(index)

	name:=[]string{"Aryan", "Ishika", "Anup", "Nayan", "Aditi", "Patel"}
	// sorting the string slices
	sort.Strings(name)
	fmt.Println(name)

	// searching a string in string slices, if the element not presnt the slices then it returns the length of slices which is not a valid index
	fmt.Println(sort.SearchStrings(name, "Aryan"))
}