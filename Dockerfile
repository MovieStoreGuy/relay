FROM golang:1.10.3-alpine AS Builder

WORKDIR /go/src/github.com/MovieStoreGuy/relay
COPY . .

RUN apk --no-cache add git && \
    go get -v -u ./... && \
    CGO_ENABLED=0 GOOS=linux go build -o /relay -ldflags="-s -w" 

FROM scratch

COPY --from=Builder /relay /relay

ENTRYPOINT ["/relay"]
