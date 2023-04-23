FROM golang:latest

WORKDIR /app
COPY . .

# Build the shared implementation library and run unit tests
WORKDIR impl
RUN rm -f go.mod
RUN go mod init "gtab/impl"
RUN go mod tidy
RUN go test
RUN GOOS=linux go build

# Build the gin-based webserver to respond to REST API requests
WORKDIR ../cmd/rest
RUN rm -f go.mod
RUN go mod init "gtab/rest"
RUN go mod edit -replace "gtab/impl"=../../impl
RUN go mod tidy
RUN go test
RUN GOOS=linux go build
RUN cp -p rest /usr/local/bin/gtab_rest_service

# Launch the webservice; e.g. use "docker run -p 80:7777/tcp ..." to map 7777 to 80 on the host
EXPOSE 7777/tcp
ENTRYPOINT ["gtab_rest_service"]

