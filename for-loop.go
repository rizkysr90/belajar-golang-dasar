package main

import "fmt"

func main() {
	// While loop
	var target uint8 = 10
	for target < 15 {
		fmt.Println("Hello", target)
		target++
	}

	// for loop
	for counter := 1; counter < 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	// for range
	arr := []int{1, 2, 3, 4, 5}

	for index, data := range arr {
		fmt.Println("Index ke", index, "dan data", data)
	}

	m := map[uint8]uint8{
		1: 11,
		2: 22,
	}

	for index, data := range m {
		fmt.Println("Index ke", index, "dan data", data)
	}
}
