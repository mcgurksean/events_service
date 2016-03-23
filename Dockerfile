FROM golang:latest

ADD . /go/src/github.com/mcgurksean/events_service

RUN go install github.com/mcgurksean/events_service

ENTRYPOINT /go/bin/events_service

EXPOSE 8080