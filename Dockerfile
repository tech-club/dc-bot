FROM golang:1.15

# Set the Current Working Directory inside the container
WORKDIR /src

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# Build the package
RUN go build

# Run the executable
CMD ["./dc-bot", "start"]
