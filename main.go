package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "csv file")
	limit := flag.Int("limit", 30, "time limit for the quiz")
	flag.Parse()
	_ = limit
	// Open the CSV file containing quiz questions and answers
	file, err := os.Open(*csvFileName)
	if err != nil {
		// Handle error (e.g., file not found or permission issues)
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	record, err := reader.ReadAll()
	if err != nil {
		// Handle CSV parsing errors (e.g., malformed data)
		fmt.Println("Error reading CSV:", err)
		return
	}
	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	correct := 0
	l := len(record)
	// Iterate over all the CSV enteries
	for i := 0; i < l; i++ {
		select {
		// Shutdown the program when the timer fires off
		case <-timer.C:
			fmt.Print("Time have run out\n")
			fmt.Printf("You have scored %d out of %d\n", correct, l)
			return
		default:
			fmt.Println(record[i][0])
			var ans string
			fmt.Scan(&ans)
			// Increment the counter for each correct answer
			if ans == record[i][1] {
				correct++
			}

		}

	}
	fmt.Printf("You have scored %d out of %d\n", correct, l)

}
