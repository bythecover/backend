# By The Cover
*author: Grant S Ralls*

# Stack
- Go 1.23
- Templ
- HTMX
- Postgres
- Host: Railway

# Dev environment setup
## Prerequisites:
- [Air](https://github.com/air-verse/air)
- [Templ](https://templ.guide/quick-start/installation)
- [Node Version Manager](https://github.com/nvm-sh/nvm)
- [Go 1.23](https://go.dev/dl/)
- Optional: [Docker](https://www.docker.com/get-started/)

## Setup Steps *(Run commands in the project root director unless specified otherwise)*
1. `nvm use` - sets the right node version
2. `npm i` - installs node dev dependencies (mainly tailwind)
3. `go mod tidy` - installs and cleans go dependencies
4. `air .` - runs the dev server, 8080 for the basic server, 8081 for hot reloading

## Building a production image and running it
While this is how a production image is built, this will fail locally. Railway injects Environment variables into the instance. I'm not doing that here. If you need to understand how a binary is made, take a look at the Dockerfile.
1. `docker build -t bythecover .`
2. `docker run --name bythecover -t bythecover`

