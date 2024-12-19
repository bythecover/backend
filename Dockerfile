FROM golang:1.22-alpine
WORKDIR /user/src/app
COPY . ./
RUN go mod download && go mod verify
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN go build
EXPOSE 8080
CMD ["./backend"]
