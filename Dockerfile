FROM golang:1.21-alpine AS base
RUN apk add --no-cache make
RUN apk add curl
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.2
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

FROM base AS builder
COPY . ./
RUN make generate-swag-doc
RUN mkdir ./build
RUN find . -name "*.json" -exec cp --parents '{}' ./build \;
RUN go build -o /app/build/poc-template cmd/app/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/build /app
CMD [ "/app/poc-template" ]