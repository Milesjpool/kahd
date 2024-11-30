FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN mkdir build
ENV CGO_ENABLED=0
RUN go build -o build ./...

FROM builder AS e2e-builder
RUN go build -o build ./cmd/e2e-tests

FROM alpine:latest AS e2e-tests
COPY --from=e2e-builder /app/build/e2e-tests ./main
EXPOSE 8080
CMD ["./main"]

FROM alpine:latest AS api-server
COPY --from=builder /app/build/api-server ./main
EXPOSE 8080
CMD ["./main"]

