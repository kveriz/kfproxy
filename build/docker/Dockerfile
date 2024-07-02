ARG GO_VERSION="1.22.2"
ARG PLATFORM="linux/amd64"
ARG ARCH="amd64"
ARG OS="linux"

##
## Build
##
FROM --platform=${PLATFORM} alpine:edge AS build

WORKDIR /app

COPY ../.. ./
RUN apk add --no-cache --update go gcc g++ libc-dev git make build-base

RUN go env -w GO111MODULE=on && \
    CGO_ENABLED=1 GOARCH=${ARCH} GOOS=${OS} go build -tags musl -o /kfproxy ./cmd/proxy/main.go

##
## Deploy
##
FROM --platform=${PLATFORM} alpine:edge AS runtime

WORKDIR /app

COPY --from=build /kfproxy /app/kfproxy

ENTRYPOINT [ "/app/kfproxy", "-port", ":8080" ]