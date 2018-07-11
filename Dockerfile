FROM golang:alpine as builder

RUN apk --no-cache add curl git
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/github.com/neoplatonist/botManager
COPY . .

RUN dep ensure; go build ./cmd/server

FROM alpine:latest

ARG DISCORD_APP_USER
ENV DISCORD_APP_USER=$DISCORD_APP_USER

RUN apk add --no-cache tzdata
ENV TZ=Asia/Tokyo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apk --no-cache add --update ca-certificates && rm -rf /var/cache/apk/* /tmp/*

COPY --from=builder /go/src/github.com/neoplatonist/botManager/server .
CMD ./botmanager-server