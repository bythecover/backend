# Copy the source & build css
FROM node:23-alpine AS source
WORKDIR /user/src/app
COPY . /user/src/app
RUN npm i
RUN npm run build

# Generate templates & build binary
FROM golang:1.23-alpine AS generate
COPY --from=source /user/src/app /user/src/app
WORKDIR /user/src/app
RUN go build

# Copy in only the binary to make a small final image
FROM scratch
COPY --from=generate /user/src/app/backend /bin/backend
CMD ["/bin/backend"] 
