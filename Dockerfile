# docker build -t nixozach/grpc-echoserver . --platform linux/amd64,linux/arm64

# Start from the official Go image
FROM golang:tip-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod ./

# Copy the source code
COPY . .

# The go env command requires git, remove this line if removing the goproxy=direct command.
RUN apk add git

# Don't confuse go build, we are only building the server in the container.
RUN rm client.go

# Needed to build in some corporate environments.
RUN go env -w GOPROXY=direct

# Build the application
RUN go build -o grpc-echoserver

# Expose the TCP port
EXPOSE 50051/tcp

# Run the application
ENTRYPOINT ["./grpc-echoserver"]