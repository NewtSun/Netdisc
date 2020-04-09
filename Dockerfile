
#build stage 二段构建
FROM golang:alpine AS builder
ENV GOPROXY=https://goproxy.io
# WORKDIR /go/src/app
ADD . /Netdisc
WORKDIR /Netdisc
# COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server

#final stage
FROM alpine:latest
ENV REDIS_ADDR=""
ENV REDIS_PW=""
ENV REDIS_DB=""
ENV MysqlDSN=""
ENV GIN_MODE="release"
ENV PORT=3000

# RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
RUN apk update && \
    apk add ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf && \
    mkdir -p /www/conf
WORKDIR /www

COPY --from=builder /Netdisc/api_server /usr/bin/api_server
ADD ./conf /www/conf

RUN chmod +x /usr/bin/api_server

ENTRYPOINT ["api_server"]
