FROM golang:1.19.5-alpine3.17 as build

ENV GOPATH /go

WORKDIR $GOPATH/src/app

COPY . .

RUN go mod init

RUN go mod download

RUN go mod tidy

RUN go build -ldflags "-s -w" -o apiserver

RUN chmod +x apiserver

FROM alpine:3.17

COPY --from=build go/src/app/apiserver /

EXPOSE 8080

ENTRYPOINT [ "/apiserver" ]