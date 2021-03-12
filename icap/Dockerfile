FROM ubuntu:20.04

RUN apt-get -y update \
 && apt-get -y install --no-install-recommends netcat \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

EXPOSE 1344

ENTRYPOINT ["nc", "-l", "1344"]

