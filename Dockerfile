# Stage 1: Build
FROM golang:1.23-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main ./main.go

# Stage 2: Runtime
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the compiled binary from the build stage
COPY --from=build /app/main .

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
