FROM golang:1.23-alpine as builder
WORKDIR /app
COPY go.mod ./
COPY *.go ./
RUN go build -o server .

FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]