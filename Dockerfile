# Use Golang base image
FROM golang:1.20

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy all the code into the container
COPY . .

# Download dependencies
RUN go mod tidy

# Build the Go app
RUN go build -o sanity-check .

# Command to run the executable
# CMD ["./sanity-check"]
CMD ["sh", "-c", "sleep 10 && ./sanity-check"]

