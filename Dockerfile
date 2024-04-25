
FROM golang:1.17-alpine as builder


WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./


RUN go mod download


COPY . .

# Build the Go app
RUN go build -o main .


FROM alpine:latest


WORKDIR /root/


COPY --from=builder /app/main .


EXPOSE 8080

the executable
CMD ["./main"]
