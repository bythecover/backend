FROM golang:1.20-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY internal/ ./internal
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o main .
EXPOSE 8080
CMD ["./main"]
