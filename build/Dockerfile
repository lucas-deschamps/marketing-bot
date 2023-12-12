FROM golang:1.21.5-bookworm AS builder

EXPOSE 3456

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o marketing-bot

FROM gcr.io/distroless/base-debian10

COPY --from=builder /app/marketing-bot /

CMD ["/marketing-bot"]
