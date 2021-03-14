#!/bin/sh
ip=`host icap | cut -d " " -f 4`
echo "$ip icap" >> /etc/hosts
/usr/local/squid/sbin/squid -d 1 -f /usr/local/squid/etc/squid.conf -N

