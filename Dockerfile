# builder image
FROM golang:1.21.5 as builder

ARG VERSION=dev

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main -ldflags=-X=main.version=${VERSION} main.go


# generate clean, final image for end users
FROM alpine:latest
COPY --from=builder /app/main .
CMD ["./main"]