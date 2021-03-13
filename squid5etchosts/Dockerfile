FROM hsmtkk/squidicap:squid5

RUN apt-get -y update \
 && apt-get -y install --no-install-recommends dnsutils \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

COPY squid5etchosts/entrypoint.sh /opt/entrypoint.sh

EXPOSE 3128

ENTRYPOINT ["/opt/entrypoint.sh"]

