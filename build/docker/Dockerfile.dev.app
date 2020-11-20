FROM golang:1.15.5-alpine3.12 AS build

RUN apk add --no-cache ca-certificates

WORKDIR /go/src/github.com/purusah/cxtnxbr
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app cmd/cxtnxbr/main.go

FROM alpine:3.12

COPY --from=build /app /app
COPY assets /assets

RUN chmod +x /app
