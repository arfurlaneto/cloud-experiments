FROM golang:1.21.6 as builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server

#FROM scratch
FROM golang:1.21.6
ARG APP_VERSION
ENV APP_VERSION ${APP_VERSION}
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/server .
ENTRYPOINT ["./server"]