FROM golang:1.16 AS builder

RUN mkdir -p /go/src/github.com/hsmtkk/squidicap

WORKDIR /go/src/github.com/hsmtkk/squidicap

COPY embip embip
COPY getip getip
COPY go.mod .
COPY go.sum .
COPY main.go .

ENV CGO_ENABLED=0

RUN go build -o /opt/embip.bin

FROM hsmtkk/squidicap:squid5

COPY --from=builder /opt/embip.bin /opt/embip.bin
COPY squid5rawip/entrypoint.sh /opt/entrypoint.sh

RUN apt-get -y update \
 && apt-get -y install --no-install-recommends less strace tcpdump \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["/opt/entrypoint.sh"]

