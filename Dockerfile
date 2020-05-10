FROM golang:1.14.1-alpine

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8080
CMD ["webmotors-search-go"]