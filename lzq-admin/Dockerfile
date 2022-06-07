FROM golang:1.17.10 as builder

ENV WORKDIR /app

WORKDIR $WORKDIR
COPY . .

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
RUN go build -o lzq-admin .

ENV TZ Asia/Shanghai

EXPOSE 80
ENTRYPOINT ["./lzq-admin"]

