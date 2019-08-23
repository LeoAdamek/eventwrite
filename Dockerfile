FROM golang:latest AS builder

ENV CGO_ENABLED 0
RUN mkdir /usr/src/app
WORKDIR /usr/src/app

COPY . .
RUN go build -v -o /usr/local/bin/eventwrite

FROM alpine:3.8
LABEL maintainer "Leo Adamek <leo.adamek@mrzen.com>"

RUN apk add --no-cache ca-certificates

RUN adduser -SH eventwrite
EXPOSE 8080

COPY --from=builder /usr/local/bin/eventwrite /usr/local/bin/eventwrite

USER eventwrite
CMD ["/usr/local/bin/eventwrite"]