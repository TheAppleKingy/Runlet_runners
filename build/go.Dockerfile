FROM golang:1.24.5-alpine 

ENV GOCACHE=/tmp/.cache
ENV GOTMPDIR=/tmp
ENV LANG=go

RUN addgroup -S runner && adduser -S runner -G runner
USER runner
WORKDIR /home/runner

COPY --from=runner_store /server /usr/local/bin/server

EXPOSE 50051

ENTRYPOINT ["/usr/local/bin/server"]
