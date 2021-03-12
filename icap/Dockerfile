FROM golang:1.16 AS builder

RUN mkdir -p /go/src/github.com/hsmtkk/squidicap

WORKDIR /go/src/github.com/hsmtkk/squidicap

COPY icap/go.mod .
COPY icap/go.sum .
COPY icap/redirect.go .

ENV CGO_ENABLED=0

RUN go build -o /opt/redirect.bin

FROM alpine:3.13.2

COPY --from=builder /opt/redirect.bin /opt/redirect.bin

EXPOSE 1344

ENTRYPOINT ["/opt/redirect.bin"]

