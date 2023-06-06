FROM golang:alpine AS builder

WORKDIR /go/src/github.com/axelwong/CqepcAuto
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add tzdata  && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

COPY --from=builder /go/src/github.com/axelwong/CqepcAuto/server ./server
COPY --from=builder /go/src/github.com/axelwong/CqepcAuto/view ./view
COPY --from=builder /go/src/github.com/axelwong/CqepcAuto/static ./static

VOLUME /app/db
VOLUME /app/log

EXPOSE 10000
ENTRYPOINT ["./server"]