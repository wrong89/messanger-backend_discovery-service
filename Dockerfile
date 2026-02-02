FROM golang:alpine AS builder

WORKDIR /app

ADD go.mod .

COPY . .

CMD [ "go", "run", "/app/main.go" ]

# RUN go build -o main main.go

# FROM alpine

# WORKDIR /app

# COPY --from=builder /app/main /app/main

# EXPOSE 9091

# CMD ["./main"]