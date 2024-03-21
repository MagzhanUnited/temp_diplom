FROM golang:latest

ENV GO111MODULE=on
ENV PORT=2002
WORKDIR /app
COPY go.mod /app
COPY go.sum /app

# Build the Go application
COPY . /app
RUN go build -o main .

# Expose the port on which your application will run
EXPOSE $PORT

# Command to run the executable
CMD ["./main"]
