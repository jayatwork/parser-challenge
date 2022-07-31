package main

import (
        "encoding/csv"
        "fmt"
        "log"
        "os"
		"time"
		"strconv"
)

const (
	jFound string = "jeff22"
    dFormat string = "2020-04-15"
)

// User object for handling user event data
type Event struct {
        Timestamp string `json:"timestamp"`	
		Username string `json:"username"`
		Operation string `json:"operation"`
		Size int `json:"size"`
}

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
		count := 0
		jevents := 0
		

		// Unique users array
		users := make([]string, len(data))

        for _, col := range data {
				size , _ := strconv.Atoi(col[3])


				timeStampString := col[0]
				layOut := "Jan 2, 2006"
				timeStamp, err := time.Parse(layOut, timeStampString)
	   
				if err != nil {
						fmt.Println(err)
						os.Exit(1)
				}
	   
				hr, min, sec := timeStamp.Clock()
	   
				fmt.Printf("Clock : [%d]hour : [%d]minutes : [%d] seconds \n", hr, min, sec)
	   
				year, month, day := timeStamp.Date()
	   
				fmt.Printf("Date : [%d]year : [%d]month : [%d]day \n", year, month, day)
	   
		
	
                event := Event{Timestamp: timeStamp,
                        Username:  col[1],
                        Operation: col[2],
                        Size:      size,
                }
				
				if  size >= 50 && col[2] == "upload" {
					uploads++
					events = append(events, event)
				}
				users = append(users, col[1])
				count++

				if  col[1] == jFound {
					jevents++
				}
        }

		fmt.Println("\n\n-----------------------------------------------------------------")
		fmt.Println("\nUploads over 50kb : ", uploads)
		//filter collection by unique users
		fmt.Println("\nNumber of events captured in the server stdout : ", count)

    	fmt.Println("\nNumber of unique users : ",  len(removeDuplicates(users))-2)

		fmt.Println("\nUser jeff22 uploaded to server on April 15th, 2020 : ", jevents , " times")  //prettify the output
		fmt.Println("\n\n-----------------------------------------------------------------")
}