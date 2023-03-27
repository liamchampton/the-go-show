FROM golang:1.19.5-alpine3.17 as build

ENV GOPATH /go

WORKDIR $GOPATH/src/app

COPY . .

RUN go mod init

# Uncomment the following line if you are using the Gin or Fiber webserver
# RUN go mod download

RUN go build -ldflags "-s -w" -o website

RUN chmod +x website

FROM alpine:3.17

COPY --from=build go/src/app/index.html /
COPY --from=build go/src/app/website /

EXPOSE 8080

ENTRYPOINT [ "/website" ]