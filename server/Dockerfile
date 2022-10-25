FROM golang:1.19-alpine as builder

WORKDIR /build
COPY ./server .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o server .

FROM alpine
WORKDIR /app
COPY --from=builder /build/server .
USER 65532:65532

ENTRYPOINT ["/app/server"]

