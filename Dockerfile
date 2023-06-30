FROM golang:1.19-alpine
# Install gcc tools, required for unit tests
RUN apk add build-base

MAINTAINER mjkgit@gmail.com

ENV GIN_MODE=release
ENV PORT=7777

WORKDIR /go/src/gtab
COPY . /go/src/gtab

# Build the shared implementation library and run unit tests
WORKDIR /go/src/gtab/impl
RUN rm -f go.mod
RUN go mod init "gtab/impl"
RUN go mod tidy
RUN go test
RUN GOOS=linux go build

# Build the gin-based webserver to respond to REST API requests
WORKDIR /go/src/gtab/cmd/rest
RUN rm -f go.mod
RUN go mod init "gtab/rest"
RUN go get github.com/gin-gonic/gin
RUN go mod edit -replace "gtab/impl"=../../impl
RUN go mod tidy
RUN go test
RUN GOOS=linux go build
RUN cp -p rest /go/src/gtab/gtab_rest_service

# Launch the webservice; e.g. use "docker run -d -p 7777:7777
WORKDIR /go/src/gtab
EXPOSE $PORT
ENTRYPOINT ["./gtab_rest_service"]

