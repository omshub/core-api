# Build container
FROM golang:1.18-bullseye as builder
WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 make build

# Run container
FROM alpine:3.15.4

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/core-api /usr/local/bin/

COPY config.yml config.yml

CMD ["core-api"]
