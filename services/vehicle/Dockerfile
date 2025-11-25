# builder
FROM golang:1.25-alpine AS builder
RUN apk add --no-cache git build-base golangci-lint

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN make lint
RUN make vehicle-test
RUN make vehicle-build

# runtime
FROM alpine:3.18 AS runtime
RUN apk add --no-cache ca-certificates
COPY --from=builder /bin/vehicle /bin/vehicle
EXPOSE 8081
USER 1000:1000
ENTRYPOINT ["/bin/vehicle"]