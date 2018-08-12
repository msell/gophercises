package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

// https://github.com/gophercises/quiz
func main() {
	d := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
	ch := make(chan Result)
	go startQuiz(ctx, ch)
	r := <-ch

	fmt.Println("You got", r.NumberCorrect, "out of", r.NumberOfQuestions, "correct")
}

func startQuiz(ctx context.Context, ch chan Result) {
	result := Result{}
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(bufio.NewReader(file))

	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		result.NumberOfQuestions++
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
			result.NumberCorrect++
		}
	}
	ch <- result
}

// Result of quiz
type Result struct {
	NumberOfQuestions int
	NumberCorrect     int
}
