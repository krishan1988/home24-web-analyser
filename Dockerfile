FROM golang:1.19.2-alpine3.16

RUN mkdir /opt/home24-web-analyser

ADD . /opt/home24-web-analyser

WORKDIR /opt/home24-web-analyser

RUN go build -o bin/home24-web-analyser cmd/web/web.go