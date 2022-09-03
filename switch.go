package main

import "fmt"

func main() {
	var number uint8 = 100

	switch number {
	case 100:
		fmt.Println("Perfect Score")
	case 85:
		fmt.Println("So good")
	case 70:
		fmt.Println("It's okay not to be okay")
	default:
		fmt.Println("Anda belum melaksanakan ujian")
	}
}
