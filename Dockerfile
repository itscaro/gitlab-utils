# Build stage
FROM golang:1.17 as builder

ARG GIT_COMMIT
ARG VERSION

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build --ldflags "-s -w \
        -X github.com/itscaro/gitlab-utils/utils.GitCommit=${GIT_COMMIT} \
        -X github.com/itscaro/gitlab-utils/utils.Version=${VERSION}" \
        -a -installsuffix cgo -o build/cli

# Release stage
FROM alpine:3.12

RUN apk --no-cache add ca-certificates git

WORKDIR /app

COPY --from=builder /app/build/cli /app/cli

ENTRYPOINT ["/app/cli"]
