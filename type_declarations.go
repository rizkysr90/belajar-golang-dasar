package main

import "fmt"

func main() {
	type NoKTP string

	var ktpRizki NoKTP = "08129384745"
	fmt.Println(ktpRizki)
	fmt.Println(NoKTP("awkwkwkkw"))
}
