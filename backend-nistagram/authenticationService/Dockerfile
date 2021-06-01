FROM golang:alpine
RUN mkdir /app
Add . /app/
WORKDIR /app
RUN go build -o main .
CMD ["./main"]
EXPOSE 8081