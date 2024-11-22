package main

import (
	"QuizGame/utils"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func gameInitializer(data [][]string, count *int) {
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

func backgroundTimer(quit chan interface{}, timer *int, count *int) {
	go func() {
		defer close(quit)
		time.Sleep(time.Duration(*timer) * time.Second)
		quit <- "done"
	}()
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Number of correct answers: ", *count)
				os.Exit(0)
			}
		}
	}()
}

func main() {
	count := 0
	timer := flag.Int("timer", 30, "Timer of the game")
	flag.Parse()
	quit := make(chan interface{})
	go backgroundTimer(quit, timer, &count)
	data, err := utils.LoadFile()
	if err != nil {
		log.Fatal(err)
	}
	gameInitializer(data, &count)
}
