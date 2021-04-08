## Telling to use Docker's golang ready image
#FROM golang:onbuild
#EXPOSE 8080
#
#RUN apt-get update && apt-get install -y mysql-client
#RUN go get -u github.com/go-sql-driver/mysql
#RUN go get -u github.com/thegenuinegourav/WalletService
## ADD . /home
#
## WORKDIR /home
#
## RUN \
##         apk add --no-cache bash git openssh && \
##         go get -u github.com/minio/minio-go
#
#
## CMD [ "go", "run", "main.go" ]

# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /go/src/wallet-service

# Expose port 8080 to the outside world
EXPOSE 8080

RUN apt-get update && apt-get install -y mariadb-client

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]