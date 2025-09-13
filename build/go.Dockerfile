FROM golang:1.24.5-alpine 

ENV GOCACHE=/tmp/.cache
ENV GOTMPDIR=/tmp
ENV LANG=Go
ENV LANGS_CONF_PATH=/home/runner/languages.yaml

RUN addgroup -S runner && adduser -S runner -G runner
USER runner
WORKDIR /home/runner
COPY languages.yaml .

COPY --from=runner_store /server /usr/local/bin/server

EXPOSE 50051

ENTRYPOINT ["/usr/local/bin/server"]
