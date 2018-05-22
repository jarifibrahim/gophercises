package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
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
	total := len(records)
	correct := 0
	// Start the quiz
	go func() {
		for index, record := range records {
			fmt.Printf("Problem #%d: %s = ", index+1, record[0])
			fmt.Scanln(&input)
			if input == record[1] {
				correct = correct + 1
			}
		}
		end(correct, total)
	}()

	// Wait for timeout
	<-time.After(time.Duration(*tlimit) * time.Second)
	end(correct, total)
}

func end(correct int, total int) {
	fmt.Printf("\nYou scored %d out of %d.\n", correct, total)
	os.Exit(0)
}
