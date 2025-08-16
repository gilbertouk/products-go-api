FROM golang:1.24.5

# Set working directory
WORKDIR /go/src/app

# Copy the source code
COPY . .

# EXPOSE the port
EXPOSE 8000

# Build the application
RUN go build -o main cmd/main.go

# Run the application (executable)
CMD ["./main"]
