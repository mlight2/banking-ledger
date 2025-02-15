FROM golang:1.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o banking-ledger ./cmd/api/main.go
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/banking-ledger .
EXPOSE 8080
CMD ["./banking-ledger"]
