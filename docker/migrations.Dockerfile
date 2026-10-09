FROM alpine:3.20

ARG GOOSE_VERSION=v3.14.0
ARG TARGETARCH

RUN case "$TARGETARCH" in \
      amd64) ARCH=x86_64 ;; \
      arm64) ARCH=arm64 ;; \
      *) echo "unsupported arch: $TARGETARCH" && exit 1 ;; \
    esac && \
    wget -qO /bin/goose "https://github.com/pressly/goose/releases/download/${GOOSE_VERSION}/goose_linux_${ARCH}" && \
    chmod +x /bin/goose

WORKDIR /app

COPY migrations migrations/
COPY --chmod=755 migrations.sh ./

ENTRYPOINT [ "migrations.sh" ]