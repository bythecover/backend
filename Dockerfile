# Copy the source & build css
FROM node:23-alpine AS source
WORKDIR /user/src/app
COPY . /user/src/app
RUN npm i
RUN npm run build

# Generate template files
FROM ghcr.io/a-h/templ:v0.3.819 AS generate
WORKDIR /user/src/app
COPY --chown=65532:65532 --from=source /user/src/app /user/src/app
RUN ["templ", "generate"]

# Generate templates & build binary
FROM golang:1.23-alpine AS build
COPY --from=generate /user/src/app /user/src/app
WORKDIR /user/src/app
RUN go mod download
RUN go build
RUN apk add --no-cache ca-certificates

# Copy in only the binary to make a small final image
FROM scratch
# Copy the CA certificates from the builder image
COPY --from=build /etc/ssl/certs /etc/ssl/certs
COPY --from=build /usr/share/ca-certificates /usr/share/ca-certificates
COPY --from=build /user/src/app/backend /bin/backend
CMD ["/bin/backend"] 
