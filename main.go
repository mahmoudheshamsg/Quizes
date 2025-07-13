package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
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

	correct := 0
	l := len(record)
	// Iterate over all the CSV enteries
	for i := 0; i < l; i++ {
		fmt.Println(record[i][0])
		rans, e := strconv.Atoi(record[i][1])
		if e != nil {

		}
		var ans int
		fmt.Scan(&ans)
		// Increment the counter for each correct answer
		if ans == rans {
			correct++
		}
	}
	fmt.Println(correct)

}
