FROM golang:1.11.10-stretch
RUN mkdir -p /go/src/github.com/gnydick/omdbapi
WORKDIR /go/src/github.com/gnydick/omdbapi
ADD . .
RUN go get
RUN go build -o main .
