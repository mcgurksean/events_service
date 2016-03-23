# golang image where workspace (GOPATH) configured at /go
FROM golang:latest

# copy the latest files to the container's workspace
ADD . /go/src/github.com/mcgurksean/events_service

# Build the events_service command inside the container
RUN go install github.com/mcgurksean/events_service

# Run the golang-docker command when the container starts
ENTRYPOINT /go/bin/events_service

# http server listens on port 8080
EXPOSE 8080