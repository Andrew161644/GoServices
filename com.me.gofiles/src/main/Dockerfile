FROM golang:1.12-alpine

RUN apk add --no-cache git

WORKDIR /app/go-sample-app
COPY . .

RUN go get -d -v ./...
RUN go build -o ./out/go-sample-app .

# This container exposes port 8080 to the outside world
EXPOSE 8080
# Run the binary program produced by `go install`
CMD ["./out/go-sample-app"]
