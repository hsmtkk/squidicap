#!/bin/sh
while true
do
	curl -x squid4:3128 http://nginx/data
	sleep 30
	curl -x squid5:3128 http://nginx/data
	sleep 30
done

