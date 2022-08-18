FROM golang:1.18-alpine AS build-env

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN apk add --no-cache upx || \
    go version && \
    go mod download
COPY . .
RUN CGO_ENABLED=0 go build -buildvcs=false -trimpath -ldflags '-w -s' -o /go/bin/rev
RUN [ -e /usr/bin/upx ] && upx /go/bin/rev || echo
FROM scratch
COPY --from=build-env /go/bin/rev /go/bin/rev
ENTRYPOINT ["/go/bin/rev"]

