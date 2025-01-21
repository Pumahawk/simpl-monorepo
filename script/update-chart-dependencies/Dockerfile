FROM alpine:3.20

RUN apk add --no-cache --allow-untrusted --no-check-certificate \
	kubectl \
	helm \
	jq \
	yq \
	bash \
	git \
	curl

WORKDIR /app

COPY update-chart-dependencies.bash /app/update-chart-dependencies.bash

CMD /app/update-chart-dependencies.bash
