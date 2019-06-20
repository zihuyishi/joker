FROM golang:1.12.6

WORKDIR /go/src/github.com/zihuyishi/joker

ENV GO111MODULE on
ENV GOPROXY https://goproxy.io

COPY . .

RUN go install -v ./app/joker.go

CMD ["joker"]