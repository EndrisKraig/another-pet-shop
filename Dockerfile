# syntax=docker/dockerfile:1

FROM golang:1.18-alpine as builder
WORKDIR /build
COPY go.mod . 
RUN go mod tidy && go mod download
COPY . .
RUN go build -buildvcs=false -o /main

FROM alpine:3
COPY --from=builder main /bin/main
COPY --from=builder /build/.env .
ENTRYPOINT ["/bin/main"]
