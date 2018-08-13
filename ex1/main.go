package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
	"flag"
)

// https://github.com/gophercises/quiz
func main() {
	timeoutPtr := flag.Int("t", 30, "Quiz timeout seconds")
	flag.Parse()

	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bufio.NewReader(file))
	numCorrect := 0
	numTotal := 0
	ch := make(chan bool)

	go func() {
		for {
			line, err := r.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			numTotal++
			question := line[0]
			answer, err := strconv.Atoi(line[1])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(question)

			var input string
			fmt.Scanln(&input)

			num, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("only numeric answers are accepted")
			}

			if num == answer {
				numCorrect++
			}
		}
		ch <- true
	}()

	timer := time.NewTimer(time.Duration(*timeoutPtr) * time.Second)
	defer timer.Stop()

	select {
	case <-ch:
		fmt.Println("Answered all questions")
	case <-timer.C:
		fmt.Println("timeout")
	}
	fmt.Println("You got", numCorrect, "out of", numTotal, "correct")
}

type problem struct {
	question string;
	answer int;
}

func getQuizData() (problems []problem, problemCount int) {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bufio.NewReader(file))

	problems = make([]problem, 0)
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		problemCount++
		question := line[0]
		answer, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}

		problem := problem{
			question: question,
			answer: answer}

		problems = append(problems, problem)
	}

	fmt.Println(problems)
	return problems, problemCount
}
