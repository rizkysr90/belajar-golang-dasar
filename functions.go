package main

import "fmt"

func sayHello() {
	// basic funtion
	fmt.Println("Hello")
}

func sayHelloWithName(name string) {
	// function paramater
	fmt.Println("Hello", name)
}
func sayHelloWithReturn(name string, greeting string) string {
	return "Hello" + " " + name + "," + greeting
}
func sayHelloWith2Return(name string) (uint8, string) {
	if name == "Rizki" {
		return 0, "Hello " + name
	}
	return 1, "Hello " + name
}
func sumAll(numbers ...int) int {
	total := 0
	// variadic function
	for _, data := range numbers {
		total += data
	}
	return total
}
func main() {
	sayHello()
	sayHelloWithName("Rizki")
	fmt.Println(sayHelloWithReturn("Ahmad", "Selamat Pagi"))
	err, data := sayHelloWith2Return("Ahmad")
	if err == uint8(1) {
		fmt.Println("Error")
	} else {
		fmt.Println(data)
	}
	fmt.Println(sumAll(1, 2, 3, 4, 5))
}
