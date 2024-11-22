package main

import (
	"QuizGame/utils"
	"fmt"
	"log"
)

func gameInitializaer(data [][]string, count *int) {
	for _, el := range data {
		var answer int
		fmt.Println("Enter the answer: ", el[0])
		_, err := fmt.Scanf("%d", &answer)
		if err != nil {
			log.Fatal(err)
		}
		if fmt.Sprintf("%d", answer) == el[1] {
			*count++
		}
	}
}
func main() {
	count := 0
	data, err := utils.LoadFile()
	if err != nil {
		log.Fatal(err)
	}
	gameInitializaer(data, &count)
	fmt.Println("Number of correct answers: ", count)
}
