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
	"math/rand"
)

// https://github.com/gophercises/quiz
func main() {
	timeoutPtr := flag.Int("t", 30, "Quiz timeout seconds")
	randomizePtr := flag.Bool("r", false, "Randomize Questions")
	flag.Parse()

	problems, problemCount := getQuizData()
	if *randomizePtr == true {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(problems), func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}

	numCorrect := 0

	fmt.Println("Press enter to begin.")
	fmt.Scanln()

	ch := make(chan bool)
	go func() {
		for _, p := range problems {

			fmt.Println(p.question)

			var input string
			fmt.Scanln(&input)

			num, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("only numeric answers are accepted")
			}

			if num == p.answer {
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
		fmt.Println("Out of time!")
	}
	fmt.Println("You got", numCorrect, "out of", problemCount, "correct")
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

	return problems, problemCount
}
