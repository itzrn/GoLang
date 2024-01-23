package main

import(
	"fmt"
)

func main()  {
	x := 0
	fmt.Println("for loop")
	for x < 5{
		fmt.Println(x)
		x++
	}

	fmt.Println("treditional for loop")
	for i:=0; i<5; i++{
		fmt.Println(i)
	}

	fmt.Println("Loop through Slices")
	names := []string{"Aryan", "Ishika", "Anup", "Aditi", "Nayan", "Patel"}
	for i:=0 ; i<len(names); i++{
		fmt.Println(names[i])
	}

	fmt.Println("Another way to loop slices")
	for index, value := range names{ // if we won't need index we can use '_' instead
		// using value we can't change the value in the slice(value ="Tea")
		fmt.Println(index, value)
	}
}
