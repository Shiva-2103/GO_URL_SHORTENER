# Stage 1: Build the application
FROM golang:alpine AS builder

# Create a build directory
RUN mkdir /build

# Copy application files to the build directory
ADD . /build/

# Set the working directory
WORKDIR /build

# Build the Go application
RUN go build -o main .

# Stage 2: Create a minimal runtime image
FROM alpine

# Add a non-root user
RUN adduser -S -D -H -h /app appuser

# Set user to non-root
USER appuser

# Copy runtime files from the first stage
COPY --from=builder /build/main /app/

# Set the working directory
WORKDIR /app

# Expose the application port
EXPOSE 3000

# Set the command to run the application
CMD ["./main"]
