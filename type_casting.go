package main

import "fmt"

func main() {
	var nilai32 int32 = 32145
	var nilai64 int64 = int64(nilai32)

	fmt.Println(nilai64)

	myname := "Rizki"
	firstCharOfMyName := string(myname[0])

	fmt.Println(firstCharOfMyName)
}
