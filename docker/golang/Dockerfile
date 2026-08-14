FROM golang:1.26-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
# RUN go mod download

# RUN go install github.com/a-h/templ/cmd/templ@latest

COPY . .

# RUN templ generate

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/web

# Running container
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/server .

COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./server"]