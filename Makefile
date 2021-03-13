default: httpclientdock icapdock squid4dock squid5dock squid5etchostsdock squid5patched

httpclientdock:
	DOCKER_BUILDKIT=1 docker build --file httpclient/Dockerfile --tag hsmtkk/squidicap:httpclient .

icapdock:
	DOCKER_BUILDKIT=1 docker build --file icap/Dockerfile --tag hsmtkk/squidicap:icap .

squid4dock:
	DOCKER_BUILDKIT=1 docker build --file squid4/Dockerfile --tag hsmtkk/squidicap:squid4 .

squid5dock:
	DOCKER_BUILDKIT=1 docker build --file squid5/Dockerfile --tag hsmtkk/squidicap:squid5 .

squid5etchostsdock:
	DOCKER_BUILDKIT=1 docker build --file squid5etchosts/Dockerfile --tag hsmtkk/squidicap:squid5etchosts .

squid5patched:
	DOCKER_BUILDKIT=1 docker build --file squid5patched/Dockerfile --tag hsmtkk/squidicap:squid5patched .
