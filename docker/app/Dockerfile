FROM golang:1.17

ENV  GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
WORKDIR /go/app/
ADD main.go .
ADD go.mod  .
RUN go install github.com/githubnemo/CompileDaemon@latest

EXPOSE 8000
ENTRYPOINT ["CompileDaemon", "-command=./vue-game", "-exclude-dir=./ui"]
