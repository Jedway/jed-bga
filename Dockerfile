# using the official golang image
FROM golang:1.20

# set up working directory
WORKDIR /app

# copy go module files
COPY go.mod .
COPY go.sum .

# download dependencies
RUN go mod download

# copy source code
COPY . .

# build application
RUN go build -o main .

# expose port
EXPOSE 8080

# run application
CMD ["./main"]
