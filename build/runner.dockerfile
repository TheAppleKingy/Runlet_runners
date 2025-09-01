FROM golang:1.24.5-alpine AS builder

WORKDIR /runner

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM scratch AS runner_store
COPY --from=builder /runner/server /server