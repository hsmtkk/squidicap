FROM ubuntu:20.04 AS builder

ADD http://www.squid-cache.org/Versions/v5/squid-5.0.5-20210223-r4af19cc24.tar.bz2 /usr/local/src/squid-5.0.5-20210223-r4af19cc24.tar.bz2

WORKDIR /usr/local/src

RUN apt-get -y update \
 && apt-get -y install --no-install-recommends gcc g++ make \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/* \
 && tar jxf squid-5.0.5-20210223-r4af19cc24.tar.bz2

WORKDIR /usr/local/src/squid-5.0.5-20210223-r4af19cc24

RUN ./configure \
 --disable-arch-native \
 --enable-icap-client \
 --prefix=/usr/local/squid \
 && make -j8 \
 && make install

FROM ubuntu:20.04

COPY --from=builder /usr/local/squid /usr/local/squid

RUN chown -R nobody:nogroup /usr/local/squid/var

ENTRYPOINT ["/usr/local/squid/sbin/squid", "-d", "1", "-f", "/usr/local/squid/etc/squid.conf", "-N"]

