#!/bin/sh
while true
do
	echo squid4
	curl -o /dev/null -x squid4:3128 http://nginx/data
	sleep 20

	echo squid5
	curl -o /dev/null -x squid5:3128 http://nginx/data
	sleep 20

	echo squid5etchosts
	curl -o /dev/null -x squid5etchosts:3128 http://nginx/data
	sleep 20
done

