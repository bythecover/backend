# Environment Variables are set with the cloud provider.
FROM alpine:latest
RUN apk add --no-cache gcompat
WORKDIR /user/src/app
COPY ./backend /user/src/app
EXPOSE 8080
CMD ["./backend"]
