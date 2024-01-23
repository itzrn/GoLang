package main

import "fmt"

func main()  {
	// array
	var ages [3]int = [3]int{20, 25, 30}
	fmt.Println(ages, len(ages))

	// short hand array
	var ages1 = [4]int{20, 25, 30, 40}
	fmt.Println(ages1, len(ages1))

	// short hand assignment
	names := [4]string{"Aryan", "Ashmit", "Anup", "Nayan"}
	fmt.Println(names, len(names))

	fmt.Println()

	// Slices -> this is more like typical array of languages where we can't change the length, can't add element to it
	var score = []int{100, 50, 60} // here we not providing length, so it will create slice
	fmt.Println("Slices before updation ->", score)
	score[2] = 25 // this way we can change any value of the slices or array
	fmt.Println("Slices after updation ->", score, len(score))

	// we can also append item to this slices which is not possible in an array
	score = append(score, 85) // this function returns slices after appending any element to it, so it is need to be store in any variable
	fmt.Println("Slices after append ->", score, len(score))

	// slice ranges
	rangeOne := names[1:3] // inclusive, exclusive
	fmt.Println("Slice ranges", rangeOne)

	// this will add the elements till end index, [:3] -> from index 0 to 2(3 is exclusive)
	rangeTwo := names[1:] // inclusive, exclusive
	fmt.Println("Slice ranges", rangeTwo)

	// at the end of the day rangeOne and rangeTwo are slices
	rangeOne = append(rangeOne, "Aditi")
	fmt.Println(rangeOne)
}

