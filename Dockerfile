FROM golang

ADD . /go/src/search-api

RUN go get github.com/99designs/gqlgen
RUN go get github.com/go-chi/chi
RUN go get github.com/olivere/elastic/v7
RUN go get github.com/vektah/gqlparser/v2


ENTRYPOINT /go/bin/search-api

EXPOSE 8080