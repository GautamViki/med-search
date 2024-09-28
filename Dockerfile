FROM golang:1.22.2-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

# If "main.go" is in the root directory, no need to specify the file.
RUN go build -o main .

# Make the binary executable
RUN chmod +x main

EXPOSE 3009

CMD ["./main"]
