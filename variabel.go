package main

import "fmt"

func main() {
	var name string = "Rizki Susilo Ramadhan"
	fmt.Println(name)

	var my_friend = "Ahmad Tabah"
	fmt.Println(my_friend)

	my_father := "Bambang Susilo"
	fmt.Println(my_father)

	var (
		my_age   uint8
		my_money int32
	)
	my_age = 22
	my_money = 3000000

	fmt.Println(my_age)
	fmt.Println(my_money)
}
