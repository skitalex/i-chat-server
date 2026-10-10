FROM --platform=$BUILDPLATFORM golang:1.27.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ARG TARGETOS TARGETARCH

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o server ./cmd/main.go

FROM alpine:3.20

WORKDIR /app/

COPY --from=builder /app/server .

CMD ["./server"]