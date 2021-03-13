#!/bin/sh
ip=`host icap5etchosts | cut -d " " -f 4`
echo "$ip icap5etchosts" >> /etc/hosts
/usr/local/squid/sbin/squid -d 1 -f /usr/local/squid/etc/squid.conf -N

