# syntax=docker/dockerfile:1

FROM golang:1.18-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY app/*.go ./
RUN go build -o /api-server
EXPOSE 4200
CMD ["/api-server"]
