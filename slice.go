package main

import "fmt"

func main() {
	var firstSlice = []int{1, 2, 3, 4, 5}
	fmt.Println(firstSlice)

	secondSlice := make([]uint8, 3, 4)
	fmt.Println(secondSlice)

	secondSlice = append(secondSlice, 100)
	fmt.Println(secondSlice)

	thirdSlice := make([]uint8, 4, 4)
	copy(thirdSlice, secondSlice)

	fmt.Println(thirdSlice)

}
