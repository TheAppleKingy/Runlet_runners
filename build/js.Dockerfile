FROM node:20-alpine

ENV NODE_ENV=production
ENV TMPDIR=/tmp
ENV HOME=/home/runner

RUN addgroup -S runner && adduser -S runner -G runner
USER runner
WORKDIR /home/runner

COPY --from=runner_store /server /usr/local/bin/server

EXPOSE 50051

ENTRYPOINT ["/usr/local/bin/server"]
