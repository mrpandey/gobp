# syntax=docker/dockerfile:experimental
FROM golang:1.21-rc-bullseye as builder

WORKDIR /app

RUN apt-get update && apt-get --no-install-recommends install -y git ca-certificates \
    && update-ca-certificates \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

RUN curl -sSf https://atlasgo.sh -o atlasgo.sh && chmod +x atlasgo.sh && ./atlasgo.sh -y
RUN curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | bash -s -- --to /usr/local/bin

# for accessing private repos on github
RUN mkdir -p -m 0600 ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts

RUN git config --global url."git@github.com:".insteadOf "https://github.com/"

COPY justfile ./

COPY go.mod go.sum ./
RUN --mount=type=ssh go mod download

COPY src ./src

RUN just build server

FROM alpine:3.18

WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/bin/atlas /usr/local/bin/atlas
COPY migrations ./migrations
COPY --from=builder /app/build/ /usr/local/bin

ARG release
ENV RELEASE_SHA $release

CMD ["gobp_server"]