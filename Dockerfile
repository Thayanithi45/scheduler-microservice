# Stage 1: Build statically linked binary
FROM golang:1.24.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# IMPORTANT: build with static linking flags
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Use scratch (smallest possible image)
FROM scratch

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

# Expose port
EXPOSE 8080

CMD ["./main"]
