FROM golang:alpine

MAINTAINER Alexander Schramm

RUN mkdir /app
ADD main.go /app/

WORKDIR /app
RUN go build -o main .

EXPOSE 4444
CMD /app/main
