FROM golang:1.21


RUN mkdir /opt/home24-web-analyser

RUN mkdir /opt/home24-web-analyser/bin

ADD . /opt/home24-web-analyser

WORKDIR /opt/home24-web-analyser

RUN go build -o bin/home24-web-analyser cmd/web/web.go