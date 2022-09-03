package main

import (
	"fmt"
)

func main() {
	scoreExam := map[string]uint16{
		"rizki": 100,
		"ahmad": 90,
		"putri": 75,
	}

	fmt.Println(scoreExam)
	fmt.Println(len(scoreExam))
	scoreExam["gahara"] = 99

	fmt.Println(scoreExam)
	fmt.Println(len(scoreExam))

	fmt.Println(scoreExam["ya"])
	delete(scoreExam, "gahara")

	fmt.Println(scoreExam)
	fmt.Println(len(scoreExam))
}
