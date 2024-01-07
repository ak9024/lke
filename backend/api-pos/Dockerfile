FROM golang:1.21 as builder
WORKDIR /root
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o api
COPY . .
FROM alpine
WORKDIR /root
RUN apk add curl
COPY --from=builder /root/api /root/api
ENTRYPOINT [ "/root/api", "server" ]

