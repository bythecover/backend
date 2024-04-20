FROM golang:1.22-alpine
WORKDIR /user/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN go install github.com/a-h/templ/cmd/templ@latest
EXPOSE 7331
CMD ["templ", "generate", "--watch", "--open-browser=false", "--proxybind=0.0.0.0", "--proxy=http://localhost:8080", "-v", "--cmd=go run ." "&", ""]
