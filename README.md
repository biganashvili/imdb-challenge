# About Project

## How to build

## How to run
```
go run main.go -primaryTitle="Titanic" -plotFilter="^.*propaganda.*$" -maxRunTime=2
```

## Help
```
  -endYear string
        filter on endYear column
  -filePath string
        absolute path to the inflated title.basics.tsv.gz file (default "./title.basics.tsv")
  -genre string
        filter on genre column
  -genres string
        filter on genres column
  -maxApiRequests uint
        maximum number of requests to be made to omdbapi
  -maxRunTime int
        maximum run time of the application. Format is a time.Duration string see here
  -originalTitle string
        filter on originalTitle column
  -plotFilter string
        regex pattern to apply to the plot of a film retrieved from
  -primaryTitle string
        filter on primaryTitle column
  -runtimeMinutes string
        filter on runtimeMinutes column
  -startYear string
        filter on startYear column
  -titleType string
        filter on titleType column
```