FROM golang:1.21.6-alpine3.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o main cmd/main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD [ "/app/main" ]