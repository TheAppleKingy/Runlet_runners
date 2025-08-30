FROM golang:1.24.5-alpine AS builder

WORKDIR /py_runner

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the runner binary
RUN CGO_ENABLED=0 GOOS=linux go build -o py_runner ./cmd/server
# ---------- Stage 2: Minimal Python runner ----------
FROM python:3.12-slim
# Environment
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1
ENV LANG python

# Create non-root user
RUN useradd -m runner
USER runner
WORKDIR /home/runner

# Copy runner binary from build stage
COPY --from=builder /py_runner/py_runner /usr/local/bin/py_runner

# Expose gRPC port
EXPOSE 50051

# Entrypoint = Go gRPC runner
ENTRYPOINT ["/usr/local/bin/py_runner"]
