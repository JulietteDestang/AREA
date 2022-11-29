# Specify the base image for the go app.
FROM golang:alpine as builder

# Specify that we now need to execute any commands in this directory.
WORKDIR /go/src/github.com/go-mysql
# Copy everything from this project into the filesystem of the container.
COPY . .
# Obtain the package needed to run redis commands. Alternatively use GO Modules.
RUN go get -u github.com/go-sql-driver/mysql
# Compile the binary exe for our app.
RUN go build -o main cmd/main/main.go
# Expose the port that will be listenned to
EXPOSE 8080
# Start the application.
CMD ["./main"]