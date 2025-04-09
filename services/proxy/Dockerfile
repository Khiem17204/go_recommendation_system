FROM golang:1.21-alpine

WORKDIR /app
COPY . .
RUN go build -o proxy main.go

EXPOSE 8001
CMD ["./proxy"]
