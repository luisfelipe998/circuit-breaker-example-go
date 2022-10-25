FROM golang:1.19-alpine as builder

WORKDIR /build
COPY ./client .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o client .

FROM alpine
WORKDIR /app
COPY --from=builder /build/client .
USER 65532:65532

ENTRYPOINT ["/app/client"]

