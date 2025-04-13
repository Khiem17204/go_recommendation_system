# Use the official Golang image to create a build artifact.
FROM golang:1.22.1-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY services/crawler/go.mod services/crawler/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy only the necessary source files
COPY services/crawler/ ./services/crawler/
COPY libs/utils/database/ ./libs/utils/database/
COPY libs/database/sqlc/ ./libs/database/sqlc/
COPY libs/utils/class/ ./libs/utils/class/

# Set the Current Working Directory inside the container
WORKDIR /app/services/crawler

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/crawler

# Start a new stage from scratch
FROM alpine:3.14

WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/crawler .

# Command to run the executable
ENTRYPOINT ["./crawler"]