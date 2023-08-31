FROM golang:1.19

RUN go version

ENV GOPATH=/

COPY ./ ./

RUN go build -o todo-app ./cmd/apiserver/main.go

CMD ["./todo-app"]
