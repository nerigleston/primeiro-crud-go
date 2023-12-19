# Build Stage
FROM golang:1.19 AS BUILDER

WORKDIR /app
COPY src src
COPY docs docs
COPY go.mod go.mod
COPY go.sum go.sum
COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
  go build -o crud-go .

# Run Stage
FROM golang:1.19-alpine3.15 as runner

RUN adduser -D neri

WORKDIR /app

COPY --from=BUILDER /app/crud-go /app/crud-go

RUN chown -R neri:neri /app
RUN chmod +x /app/crud-go

EXPOSE 8080

USER neri

CMD ["./crud-go"]
