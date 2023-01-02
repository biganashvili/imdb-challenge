package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/biganashvili/imdb-challenge/models"
	"github.com/biganashvili/imdb-challenge/queue"
)

const NUMBER_OF_GORUTINES = 10

// var lines []string
var tasks chan string

func main() {
	tasks = make(chan string, 100)
	completedTasks := queue.CreateQueue()
	filePath := flag.String("filePath", "./title.basics.tsv", "absolute path to the inflated title.basics.tsv.gz file")

	titleType := flag.String("titleType", "", "filter on titleType column")
	primaryTitle := flag.String("primaryTitle", "", "filter on primaryTitle column")
	originalTitle := flag.String("originalTitle", "", "filter on originalTitle column")
	genre := flag.String("genre", "", "filter on genre column")
	startYear := flag.String("startYear", "", "filter on startYear column")
	endYear := flag.String("endYear", "", "filter on endYear column")
	runtimeMinutes := flag.String("runtimeMinutes", "", "filter on runtimeMinutes column")
	genres := flag.String("genres", "", "filter on genres column")

	maxApiRequests := flag.Uint("maxApiRequests", 0, "maximum number of requests to be made to omdbapi")
	maxRunTime := flag.Int("maxRunTime", 0, "maximum run time of the application. Format is a time.Duration string see here")
	// maxRequests := flag.String("maxRequests", "", "maximum number of requests to send to omdbapi")
	plotFilter := flag.String("plotFilter", "", "regex pattern to apply to the plot of a film retrieved from")
	flag.Parse()

	var apiRequests uint64

	var plotRegexp *regexp.Regexp
	var err error
	if *plotFilter != "" {
		plotRegexp, err = regexp.Compile(*plotFilter)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	filters := make(map[int]string)
	if *titleType != "" {
		filters[1] = *titleType
	}
	if *primaryTitle != "" {
		filters[2] = *primaryTitle
	}
	if *originalTitle != "" {
		filters[3] = *originalTitle
	}
	if *genre != "" {
		filters[4] = *genre
	}
	if *startYear != "" {
		filters[5] = *startYear
	}
	if *endYear != "" {
		filters[6] = *endYear
	}
	if *runtimeMinutes != "" {
		filters[7] = *runtimeMinutes
	}
	if *genres != "" {
		filters[8] = *genres
	}

	wg := sync.WaitGroup{}
	workerQuitTriggers := [NUMBER_OF_GORUTINES]chan bool{}
	shouldPrintResult := true
	for i := 0; i < NUMBER_OF_GORUTINES; i++ {
		workerQuitTriggers[i] = make(chan bool)
		wg.Add(1)
		go func(name string, quit chan bool, plotRegexp *regexp.Regexp) {
			for {
				select {
				case <-quit:
					// fmt.Println(name, "ShutDown")
					wg.Done()
					return
				case line, ok := <-tasks:
					if !ok {
						wg.Done()
						return
					}
					rowData := strings.Split(line, "\t")
					if len(rowData) == 9 {
						filterPass := true
						for k, v := range filters {
							if k == 8 {
								genres := strings.Split(rowData[k], ",")
								genrePass := false
								for _, genre := range genres {
									if genre == v {
										genrePass = true
									}
								}
								if !genrePass {
									filterPass = false
									break
								}
							} else {
								if v != rowData[k] {
									filterPass = false
									break
								}
							}

						}
						if filterPass {
							if *maxApiRequests > 0 && apiRequests >= uint64(*maxApiRequests) {
								wg.Done()
								return
							}
							atomic.AddUint64(&apiRequests, 1)
							plot, err := getPlot(rowData[0])
							if err != nil {
								//todo what to do on api error
								fmt.Println(err)
								continue
							}
							if *plotFilter == "" || plotRegexp.MatchString(plot) {
								completedTasks.Insert(models.Movie{ID: rowData[0], Title: rowData[2], Plot: plot})
							}
						}
					}
				default:
					fmt.Println("nothing to do")
				}
			}
		}(fmt.Sprintf("worker_%d", i), workerQuitTriggers[i], plotRegexp)
	}

	go readFile(*filePath, tasks)

	// ------------------------------Graceful shutdown on syscall-------------\\
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		shouldPrintResult = false
		for _, quit := range workerQuitTriggers {
			quit <- true
		}
	}()
	// ------------------------------Graceful shutdown on syscall-------------//

	// ------------------------------Graceful shutdown on maxRunTime----------\\
	if *maxRunTime != 0 {
		go func() {
			time.Sleep(time.Duration(*maxRunTime) * time.Second)
			for _, quit := range workerQuitTriggers {
				quit <- true
			}
		}()
	}
	// ------------------------------Graceful shutdown on maxRunTime----------//
	wg.Wait()
	// ------------------------------Print Result-----------------------------\\
	if shouldPrintResult {
		fmt.Printf("%-12s|   %-17s   |   %s\n", "IMDB_ID", "Title", "Plot")
		for i := 0; i < len(completedTasks.List()); i++ {
			movie := completedTasks.List()[i]
			fmt.Printf("%-12s|   %-17s   |   %s\n", movie.ID, movie.Title, movie.Plot)
		}
	}
	// ------------------------------Print Result-----------------------------//
}

func readFile(filePath string, tasks chan string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		tasks <- scanner.Text()
	}
	close(tasks)
	file.Close()
	return nil
}

func getPlot(ID string) (string, error) {
	var response models.Response
	// resp, err := http.Get("http://www.omdbapi.com/?apikey=2e291c6a&i=" + ID)
	resp, err := http.Get("http://www.omdbapi.com/?apikey=b74ee860&i=" + ID)
	if err != nil {
		return "", err
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status_code:%d, body:%s", resp.StatusCode, string(body))
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}
	return response.Plot, nil
}
