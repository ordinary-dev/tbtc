FROM golang:1.19-alpine AS builder

WORKDIR /app
COPY go.mod main.go config.go parser.go ./

RUN go build -o tce .

FROM alpine:3.17

COPY --from=builder /app/tce /tce

ENTRYPOINT ["/tce"]
