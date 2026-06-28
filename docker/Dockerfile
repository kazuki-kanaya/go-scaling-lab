# build stage
FROM golang:1.26 AS builder
WORKDIR /app
COPY go.mod ./
COPY cmd ./cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/api ./cmd/api

# runtime stage
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /out/api /app/api
EXPOSE 8080
CMD ["/app/api"]
