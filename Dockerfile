# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

ADD ./ ./

RUN go mod download

RUN go build -o /imdb-challenge

CMD [ '/imdb-challenge -primaryTitle="Titanic" -genres="Drama" -plotFilter="^.*propaganda.*$" -maxRunTime=5 ' ]