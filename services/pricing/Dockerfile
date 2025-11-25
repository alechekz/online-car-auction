# builder
FROM golang:1.25-alpine AS builder
RUN apk add --no-cache git build-base golangci-lint
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN make lint
RUN make pricing-test
RUN make pricing-build

# runtime
FROM alpine:3.18 AS runtime
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/bin/grpcurl /usr/local/bin/grpcurl
COPY --from=builder /bin/pricing /bin/pricing
EXPOSE 8084
USER 1000:1000
ENTRYPOINT ["/bin/pricing"]