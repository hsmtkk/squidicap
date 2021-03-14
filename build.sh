#!/bin/sh
name=$1
DOCKER_BUILDKIT=1 docker build --file $name/Dockerfile --tag hsmtkk/squidicap:$name .

