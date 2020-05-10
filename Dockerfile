FROM golang:1.14.1-alpine as builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main server.go
EXPOSE 8080

FROM scratch
WORKDIR /root
COPY --from=builder /go/src/app/main .
CMD ["./main"]