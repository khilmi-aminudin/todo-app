FROM golang:1.19-alpine AS builder

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY . .

# download Go modules and dependencies
RUN go mod tidy

# compile application
RUN go build main.go binary
 
# tells Docker that the container listens on specified network ports at runtime
EXPOSE 8080

# command to be used to execute when the image is used to start a container
CMD [ "/binary" ]