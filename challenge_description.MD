# Golang Challenge IMDB

## In a nutshell

Write a Golang program that reads a large file, in this case the IMDB database!

While reading the file apply filters passed in as program flags, then take the remaining results and query the [omdbapi](https://www.omdbapi.com/) API to get the plot of each movie.

Lastly, check if the plot matches regex provided as a flag. The output should look as follows:

```
IMDB_ID     |   Title               |   Plot
tt0000005   |   Blacksmith Scene    |   Three men hammer on an anvil and pass a bottle of beer around. 
```

## Before you start

1. Download the file: `title.basics.tsv.gz` from [IMDB](https://datasets.imdbws.com/).
2. Get an API key at [omdbapi](https://www.omdbapi.com/)

## Requirements

1. use standard libraries where possible
2. the program must accept the following flags (all the filters are exact matching, do not cater for misspelled inputs etc):

    - filePath - absolute path to the inflated `title.basics.tsv.gz` file
    - titleType - filter on `titleType` column
    - primaryTitle - filter on `primaryTitle` column
    - originalTitle - filter on `originalTitle` column
    - genre - filter on `genre` column
    - startYear - filter on `startYear` column
    - endYear - filter on `endYear` column
    - runtimeMinutes - filter on `runtimeMinutes` column
    - genres - filter on `genres` column
    - maxApiRequests - maximum number of requests to be made to [omdbapi](https://www.omdbapi.com/)
    - maxRunTime - maximum run time of the application. Format is a `time.Duration` string see [here](https://godoc.org/time#ParseDuration)
    - maxRequests - maximum number of requests to send to [omdbapi](https://www.omdbapi.com/)
    - plotFilter - regex pattern to apply to the plot of a film retrieved from [omdbapi](https://www.omdbapi.com/)

3. if the program is still running and the `maxRunTime` threshold is reached, the program must print any results and gracefully exit 
4. the program must gracefully exit with no output when stopped with SIGTERM
5. write unit tests :)
6. write a benchmark for the section of your code which filters output from the file

## How the assessment process works

Take this coding assignment and work on it for 2-3 hours.

See the `level` breakdown below. It should give you a good idea of which features are more important than others so you can use your time appropriately.

We do not expect you to get to level 6.

Once we received your attempt and are invited for another interview, we will continue where you left off and complete your current level and/or refactor.

For example:

In your attempt you completed Level 1, however, stopped at Level 2, Task 2.4.

In the following interview you will implement the remainder of the tasks and possibly refactor (this is time dependant), as well as answer some questions regarding your code, style, approach etc.

## For the candidate

For development, we suggest you use a truncated version of the file.

### Level 1

Task 1.1: Read the file with multiple goroutines. At this point there is no limit to the number of goroutines you can use.

### Level 2

Task 2.1: Limited the number of goroutines reading/processing the file (you can add an additional flag to specify the number of goroutines. or hard code it).

Task 2.2: Take the data read from the file and apply the filters such as `primaryTitle`

Task 2.3: The program should exit when the `maxRunTime` is exceeded. At this point do not worry about a graceful exit or any resource cleanup.

Task 2.4: The program should exit when stopped with `SIGTERM`. At this point do not worry about a graceful exit or any resource cleanup.

Task 2.5: Write at least one unit test.

### Level 3

Task 3.1: Apply all of the remaining filters

Task 3.2: Query the [omdbapi](https://www.omdbapi.com/) with the filtered values. At this point do not limit the number of requests sent to the API.

Task 3.3: Output the responses from [omdbapi](https://www.omdbapi.com/). At this point we are not concerned with how the output looks, or if the REGEX filters have been applied.

Task 3.4: Write 2 more unit tests.

### Level 4

Task 4.1: Apply the given REGEX filter to the [omdbapi](https://www.omdbapi.com/) response

Task 4.2: Here we expect correctly formatted output.

Task 4.3: Your unit test coverage should be getting close to 50%

### Level 5


Task 5.1: When your program exceeds the `maxRunTime`, it should gracefully exit. All resources should be released and any running goroutines should be stopped. Any results (if available should be printed)

Task 5.2: When your program receives the `SIGTERM` signal, it should gracefully exit. All resources should be released and any running goroutines should be stopped. No results (if any) should be printed. A simple message such as "program exiting" is acceptable. 

Task 5.3: You program should be able to handle any rate limiting from [omdbapi](https://www.omdbapi.com/). It must not panic or error. The preferred behaviour is to fast exit, but not required.

Task 5.4: The required benchmark is provided.

Task 5.5: Your unit test coverage should be over 50%

### Level 6

Task 6.1: The program implements the `maxRequests` flag and fast exits when the limit is reached.

Task 6.2: Your unit test coverage should be close to 100%.

Task 6.3: You have provided everything under quality of life.

## Internal ranking

### Level 1

* Read the file and process the output in multiple goroutines (no limit on number of go routines)
* Program does not adhere to `maxRunTime`
* Program does not adhere to `SIGTERM`
* Program does not query [omdbapi](https://www.omdbapi.com/)
* Program does not filter the file
* Program has no output
* 0% unit test coverage
* requested benchmark not provided
* maxRequests not implemented

### Level 2

* Read the file and process the output in multiple goroutines (limited number of go routines)
* Program adheres to `maxRunTime` ungracefully
* Program adheres to `SIGTERM` ungracefully
* Program does not query [omdbapi](https://www.omdbapi.com/)
* Program filters the file
* Program has no output
* 0-10% unit test coverage
* requested benchmark not provided
* maxRequests not implemented

### Level 3

* Read the file and process the output in multiple goroutines (limited number of go routines)
* Program adheres to `maxRunTime` ungracefully
* Program adheres to `SIGTERM` ungracefully
* Program queries [omdbapi](https://www.omdbapi.com/), but does not apply regex filter
* Program filters the file
* Program has output
* 10-20% unit test coverage
* requested benchmark not provided
* maxRequests not implemented

### Level 4

* Read the file and process the output in multiple goroutines (limited number of go routines)
* Program adheres to `maxRunTime` ungracefully
* Program adheres to `SIGTERM` ungracefully
* Program queries [omdbapi](https://www.omdbapi.com/), applies regex filter
* Program filters the file
* Program has correctly formatted output
* 20-50% unit test coverage
* requested benchmark not provided
* maxRequests not implemented

### Level 5

* Read the file and process the output in multiple goroutines (limited number of go routines)
* Program adheres to `maxRunTime` with graceful exit
* Program adheres to `SIGTERM` with graceful exit
* Program queries [omdbapi](https://www.omdbapi.com/), applies regex filter
* Program filters the file
* Program has correctly formatted output
* Program handles rate limiting from [omdbapi](https://www.omdbapi.com/)
* 50-100% unit test coverage
* requested benchmark provided
* maxRequests not implemented

### Level 6

* Read the file and process the output in multiple goroutines (limited number of go routines)
* Program adheres to `maxRunTime` with graceful exit
* Program adheres to `SIGTERM` with graceful exit
* Program queries [omdbapi](https://www.omdbapi.com/), applies regex filter
* Program filters the file
* Program has correctly formatted output
* Program handles rate limiting from [omdbapi](https://www.omdbapi.com/)
* 100% unit test coverage
* requested benchmark provided
* maxRequests implemented

## Quality of life

* Provided Dockerfile
* Provided Readme
* Provided profiling with `pprof`