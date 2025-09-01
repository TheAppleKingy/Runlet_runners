FROM python:3.12-alpine

ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1
ENV LANG python

RUN addgroup -S runner && adduser -S runner -G runner
USER runner
WORKDIR /home/runner

COPY --from=runner_store /server /usr/local/bin/server

EXPOSE 50051

ENTRYPOINT ["/usr/local/bin/server"]
