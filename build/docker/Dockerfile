ARG GO_VERSION="1.22.2"
ARG PLATFORM="linux"
ARG OS="arm64"
FROM --platform=${PLATFORM} golang:${GO_VERSION}-alpine AS build
WORKDIR /app

COPY ../.. ./

# build app
RUN apk add --no-cache ca-certificates

RUN go mod download

RUN go env -w GO111MODULE=on && \
    CGO_ENABLED=0 GOARCH=${OS} GOOS=${PLATFORM} go build -o /kfproxy ./cmd/proxy/main.go

##
## Deploy
##
FROM --platform=${PLATFORM} scratch

WORKDIR /

COPY --from=build /kfproxy /kfproxy
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT /kfproxy -port :8080