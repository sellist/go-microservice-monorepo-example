FROM golang:alpine
WORKDIR /firstMicroservice
ADD . .
RUN go mod tidy
RUN go build main.go
EXPOSE 8080
ENTRYPOINT ["./main"]