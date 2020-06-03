FROM golang:1.14.4-buster

COPY . /app/ 

WORKDIR /app/server

RUN go build .

EXPOSE 5000

ENTRYPOINT ["./server"]
