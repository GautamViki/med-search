# Specify the base image for the Go app
FROM golang

# Specify that we now need to execute any commands in this directory
WORKDIR /medsearch

# Copy everything from this project into the filesystem of the container
COPY . .

# Copy the wait-for-it script into the container
COPY wait-for-it.sh /usr/local/bin/wait-for-it.sh

# Make sure the script is executable
RUN chmod +x /usr/local/bin/wait-for-it.sh

# Install the MySQL client
RUN apt-get update && apt-get install -y default-mysql-client

# Obtain the package needed to run MySQL commands. Alternatively, use Go Modules.
RUN go get -u github.com/go-sql-driver/mysql

# Compile the binary exe for our app
RUN go build -o main .

# Start the application
CMD ["./main"]
