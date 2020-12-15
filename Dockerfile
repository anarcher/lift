FROM golang:1.15-alpine AS builder

ENV CGO_ENABLED=0
ENV GO111MODULE=on
RUN apk add --no-cache git

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN go build -ldflags '-w -s' -o lf ./cmd/lf/

# Run on scratch
FROM alpine:3.11.3
WORKDIR /app

COPY . ./
COPY --from=builder /build/lf /app/
RUN chmod +x /app/lf

ENTRYPOINT ["/app/lf"]

