FROM golang:1.27.1 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/main.go

FROM alpine
WORKDIR /app/
COPY --from=builder /app/.env .
COPY --from=builder /app/server .
CMD ["./server"]