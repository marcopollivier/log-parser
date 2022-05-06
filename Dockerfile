FROM golang:1.18.1-alpine3.15
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o log-parser cmd/server/main.go
CMD ["/app/log-parser"]
