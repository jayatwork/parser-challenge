package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const jFound string = "jeff22"

// User object for handling user event data
type Event struct {
	Timestamp string `json:"timestamp"`
	Username  string `json:"username"`
	Operation string `json:"operation"`
	Size      int    `json:"size"`
}

//Helper utility funcs for removing duplicate users
func removeDuplicates(arr []string) []string {
	usr := map[string]bool{}
	for i := range arr {
		usr[arr[i]] = true
	}
	filtered := []string{}
	for j, _ := range usr {
		filtered = append(filtered, j)
	}
	return filtered
}

func main() {
	// find csv local to project and open
	file, err := os.Open("data/server_log.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Parse the individual comma delimited entries
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var events []Event
	uploads := 0
	count := 0
	jevents := 0

	// Unique users array
	users := make([]string, len(data))

	for _, col := range data {
		size, _ := strconv.Atoi(col[3])
		//Regex pattern matching on Apr 15 date
		match, _ := regexp.MatchString("Apr 15", col[0])

		event := Event{Timestamp: col[0],
			Username:  col[1],
			Operation: col[2],
			Size:      size,
		}

		//Conditonal check for uploads exceeding 50kb filesize and operation type
		if size >= 50 && col[2] == "upload" {
			uploads++
			events = append(events, event)
		}
		users = append(users, col[1])
		count++

		//Conditional for user jeff22 specfic operations
		if col[1] == jFound && col[2] == "upload" && match {
			jevents++
		}
	}

	fmt.Println("\n\n-----------------------------------------------------------------")
	fmt.Println("\nUploads over 50kb : ", uploads)
	fmt.Println("\nNumber of events captured in the server stdout : ", count-1)
	//filter collection by unique users
	fmt.Println("\nNumber of unique users : ", len(removeDuplicates(users))-2)
	fmt.Println("\nUser jeff22 uploaded to server on April 15th, 2020 : ", jevents, " times") 
	fmt.Println("\n\n-----------------------------------------------------------------")

	
}
