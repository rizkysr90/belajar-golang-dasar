package main

import "fmt"

func main() {
	a := "foo"
	b := "bar"
	if a == "foo" {
		fmt.Println("Oke foo")
	}
	if a == "foo" && b == "bar" {
		fmt.Println("Oke foo dan bar")
	}

	// short statement with if
	if nilai := 80; nilai > 75 {
		fmt.Println("Selamat anda lulus")
	} else {
		fmt.Println("Sorry belum lulus")
	}
}
