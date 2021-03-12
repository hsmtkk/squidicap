#!/bin/sh
while true
do
	curl -o /dev/null -x squid4:3128 http://nginx/data
	sleep 30
	curl -o /dev/null -x squid5:3128 http://nginx/data
	sleep 30
done

