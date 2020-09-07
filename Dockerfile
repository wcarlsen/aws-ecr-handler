# TODO: Make this example with our best pratices

FROM golang:1.14.4-alpine3.12 AS builder
WORKDIR /go/src/github.com/dfds/aws-ecr-handler/
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags musl -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/dfds/aws-ecr-handler/app .

ENTRYPOINT ./app


# # first stage
# FROM golang:1.14.2-stretch as build_base
# WORKDIR /src
# COPY go.mod .
# COPY go.sum .
# RUN go mod download

# # second stage
# FROM build_base AS builder
# WORKDIR /src
# COPY . .
# RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -tags musl -installsuffix cgo -o app .

# # final stage
# FROM debian:10.3-slim
# RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
# COPY --from=builder /bin/app /bin/app
# ENTRYPOINT ["/bin/app"]