FROM node:20-alpine

ENV NODE_ENV=production
ENV TMPDIR=/tmp
ENV HOME=/home/runner
ENV LANG="Java Script"
ENV LANGS_CONF_PATH=/home/runner/languages.yaml

RUN addgroup -S runner && adduser -S runner -G runner
USER runner
WORKDIR /home/runner
COPY languages.yaml .

COPY --from=runner_store /server /usr/local/bin/server

EXPOSE 50051

ENTRYPOINT ["/usr/local/bin/server"]
