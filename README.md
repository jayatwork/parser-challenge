# Coding Challenge

The intent of this coding assignment was to express proficiency in the Go Programming Language by parsing a CSV file given as input with specific requirements being:

* How many users accessed the server?
* How many uploads were larger than `50kB`?
* How many times did `jeff22` upload to the server on April 15th, 2020?
* Include a small `README.md` describing your solution and additional considerations you would make if you had more time.


## Project structure

```sh
    .
├── README.md
├── data
│   └── server_log.csv
└── main.go

```


## How It Works

1. Extract the parser challenge and change directory into <b> parser-challenge </b>:

    ```sh
    $ cd parser-challenge
    ```

2. Ensure <b>server_log.csv</b> is located under "data/" subdirectory:

    ```sh
    $ vi data/server_log.csv 
    ```

3. Run the main go file as entrypoint to program

    ```sh
    $ go run main.go
    ```


## Expected output rendered in terminal

```sh
-----------------------------------------------------------------

Uploads over 50kb :  144

Number of events captured in the server stdout :  657

Number of unique users :  6

User jeff22 uploaded to server on April 15th, 2020 :  3  times


-----------------------------------------------------------------
```