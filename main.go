package main

import (
        "encoding/csv"
        "fmt"
        "log"
        "os"
		"strconv"
)

// User object for handling user event data
type Event struct {
        Timestamp string `json:"timestamp"`
		
		Username string `json:"username"`
		Operation string `json:"operation"`
		Size string `json:"size"`
}


// - The `timestamp` is recorded in the UNIX date format.
// - The `username` is a unique identifier for the user.
// - The `operation` indicates if an `upload` or `download` occurred, no other values will appear in this column.
// - The `size` is an integer reflecting file size in `kB`.

// Your solution must be able to answer the following types of questions about the example log:
// 1. How many users accessed the server?
// 2. How many uploads were larger than `50kB`?
// 3. How many times did `jeff22` upload to the server on April 15th, 2020?


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
        for i, col := range data {
                event := Event{Timestamp: col[0],
                        Username:  col[1],
                        Operation: col[2],
                        Size:      col[3],
                }
				size , _ := strconv.Atoi(col[3])
				if  size >= 50 {
					events = append(events, event)
				}
                

            	fmt.Println("\nEvent number and detail: ", i, "\n", event)
        }

		fmt.Println(events)

}