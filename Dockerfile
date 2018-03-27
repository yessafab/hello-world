FROM golang:latest as builder
WORKDIR /go/src/github.com/yessafab/hello-world
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN GOOS=linux GOARCH=386 go build -ldflags "-s" -a -installsuffix cgo -o app app.go

FROM scratch
COPY --from=builder /go/src/github.com/yessafab/hello-world/app .
CMD ["./app"]
