FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 5001
CMD ["./main"]
