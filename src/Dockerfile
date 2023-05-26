#build stage
FROM golang:alpine AS builder
WORKDIR /src
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /server

#final stage
FROM alpine:latest
WORKDIR /web
COPY --from=builder /server .
RUN chmod +x server
CMD ["/web/server"]
EXPOSE 8080
