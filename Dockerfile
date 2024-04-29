FROM golang:1.18-alpine as builder

WORKDIR /app

# Copy the go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and other necessary directories and files into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch for a smaller final image
FROM alpine:latest  

# Install CA certificates
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/views ./views
COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["air"]
