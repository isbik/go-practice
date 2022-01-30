FROM golang:alpine

WORKDIR /files

COPY cmd/app/main.go  /files/

RUN go build -o /files/ main.go

ENTRYPOINT [ "/files/main" ]