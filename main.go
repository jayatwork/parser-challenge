package main

import (
        "encoding/csv"
        "fmt"
        "log"
        "os"
		"strconv"
		"time"
)

// User object for handling user event data
type Event struct {
        Timestamp time.Time `json:"timestamp"`
		
		Username string `json:"username"`
		Operation string `json:"operation"`
		Size int `json:"size"`
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
		uploads := 0

		//uniqueUser := "jeff22"
        for i, col := range data {
				size , _ := strconv.Atoi(col[3])
				
				t, _ := time.Parse(time.RFC3339, col[0])
                event := Event{Timestamp: t,
                        Username:  col[1],
                        Operation: col[2],
                        Size:      size,
                }
				
				if  size >= 50 && col[2] == "upload" {
					uploads++
					events = append(events, event)
				}
                
				// if  t.Truncate(24*time.Hour).Equal(t2.Truncate(24*time.Hour)) {
				// 	uploads++
				// 	events = append(events, event)
				// }


            	fmt.Println("\nEvent number and detail: ", i, "\n", event)
        }

		fmt.Println("Uploads over 50kb = ", uploads)
		//filter collection by unique users
		fmt.Println("Number of unique users accessing the server: ")

		fmt.Println("User jeff22 uploaded to server on April 15th, 2020 : ")  //prettify the output

}