#!/bin/sh
export TEMPLATE_PATH=/usr/local/squid/etc/squid.conf.tmpl
export CONF_PATH=/usr/local/squid/etc/squid.conf
export ICAP_HOST_NAME=icap
/opt/embip.bin
/usr/local/squid/sbin/squid -d 1 -f /usr/local/squid/etc/squid.conf -N

