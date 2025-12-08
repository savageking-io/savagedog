FROM golang:1.25-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN apt-get update && apt-get install -y --no-install-recommends make \
    && make linux \
    && mkdir -p /out \
    && if [ -f daemon.yaml ]; then cp daemon.yaml /out/daemon.yaml; else cp example-daemon.yaml /out/daemon.yaml; fi \
    && apt-get purge -y make \
    && rm -rf /var/lib/apt/lists/*


FROM debian:bookworm-slim AS runtime

LABEL org.opencontainers.image.source="https://github.com/savageking-io/savagedog"
LABEL org.opencontainers.image.title="savagedog"

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/bin/savagedog-linux-amd64 /usr/local/bin/savagedog

COPY --from=builder /out/daemon.yaml /etc/daemon.yaml

EXPOSE 12005

ENTRYPOINT ["/usr/local/bin/savagedog", "serve", "--config", "/etc/daemon.yaml", "--log", "trace"]