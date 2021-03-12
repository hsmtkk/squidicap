FROM ubuntu:20.04 AS builder

ADD http://www.squid-cache.org/Versions/v4/squid-4.14.tar.bz2 /usr/local/src/squid-4.14.tar.bz2

WORKDIR /usr/local/src

RUN apt-get -y update \
 && apt-get -y install --no-install-recommends gcc g++ make \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/* \
 && tar jxf squid-4.14.tar.bz2

WORKDIR /usr/local/src/squid-4.14

RUN ./configure \
 --disable-arch-native \
 --enable-icap-client \
 --prefix=/usr/local/squid \
 && make -j8 \
 && make install

FROM ubuntu:20.04

COPY --from=builder /usr/local/squid /usr/local/squid

