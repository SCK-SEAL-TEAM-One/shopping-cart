FROM golang:1.16.2 AS builder
WORKDIR /module
COPY .  .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates && \
    apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime && \
    echo "Asia/Bangkok" >  /etc/timezone && \
    apk del tzdata
WORKDIR /root/
COPY --from=builder /module/app .
ENV GIN_MODE release
EXPOSE 8000
CMD ["./app"]