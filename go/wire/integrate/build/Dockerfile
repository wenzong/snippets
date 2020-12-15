####
####        stage1
####
FROM golang:1.15.6-alpine AS build

ENV CGO_ENABLED=0 GOOS=linux GO111MODULE=on GOPROXY=https://goproxy.cn/
WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download # && go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go get

COPY . $WORKDIR
RUN go build -o /usr/local/bin/app -ldflags="-s -w" main.go && go clean -cache

####
####        stage2
####
FROM alpine:3
COPY --from=build /usr/local/bin/app /usr/local/bin/
