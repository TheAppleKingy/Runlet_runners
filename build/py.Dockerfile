FROM python:3.12-alpine

ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1
ENV LANG Python
ENV LANGS_CONF_PATH=/home/runner/languages.yaml

RUN addgroup -S runner && adduser -S runner -G runner
USER runner
WORKDIR /home/runner
COPY languages.yaml .

COPY --from=runner_store /server /usr/local/bin/server

EXPOSE 50051

ENTRYPOINT ["/usr/local/bin/server"]
