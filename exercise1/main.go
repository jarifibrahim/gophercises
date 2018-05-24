package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	csvfile := flag.String("csv", "problems.csv", "A CSV file in the format of 'question,answer'")
	tlimit := flag.Int("limit", 30, "The time limit for quiz in seconds")
	flag.Parse()

	// Open file
	file, err := ioutil.ReadFile(*csvfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := csv.NewReader(bytes.NewReader(file))
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var input string
	var done = make(chan bool)
	total := len(records)
	correct := 0
	// Start the quiz
	go func() {
		for index, record := range records {
			fmt.Printf("Problem #%d: %s = ", index+1, record[0])
			fmt.Scanln(&input)
			input = strings.TrimSpace(input)
			if strings.ToLower(input) == strings.ToLower(record[1]) {
				correct = correct + 1
			}
		}
		done <- true
	}()

	// Wait for timeout
	timeout := time.After(time.Duration(*tlimit) * time.Second)

	select {
	case <-done:
	case <-timeout:
		fmt.Println("\nTimes Up!")
	}
	fmt.Printf("\nYou scored %d out of %d.\n", correct, total)
}
