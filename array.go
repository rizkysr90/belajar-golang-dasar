package main

import "fmt"

func main() {
	var firstArray [3]uint8
	firstArray[0] = 1
	fmt.Println(firstArray)
	fmt.Println(len(firstArray))
	fmt.Println(cap(firstArray))

	// Array Literal

	var secondArray = [3]uint8{1, 2, 3}
	fmt.Println(secondArray)

	var spreadArray = [...]uint8{1, 2, 3, 4, 5}
	fmt.Println(spreadArray)

}
