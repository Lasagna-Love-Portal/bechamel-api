# syntax=docker/dockerfile:1
# Based on instructions at https://docs.docker.com/language/golang/build-images/

FROM golang:1.20 as base

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# TODO: it's laborious to add directories to copy manually,
# but maintaining a .dockerignore file and hoping people update it
# may be a worse idea - may introduce security issues if it's not maintained
# So for the time being we'll explicitly state what to copy
COPY *.go ./
COPY config ./config/
COPY documentation ./documentation/
COPY internal ./internal/
COPY model ./model/

##########
# Configuration for developer local Docker deployment
# Contains suggestions from https://dev.to/bruc3mackenzi3/debugging-go-inside-docker-using-vscode-4f67
# to enable debugging locally with VS Code via Delve
FROM base as dev
RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest
# Build the API server with gcflags that disable inlining and optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -gcflags "all=-N -l" -o /bechamel-api-localdev-server .
# Start the Delve server on port 4000
CMD [ "/go/bin/dlv", "--listen=:4000", "--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "/bechamel-api-localdev-server" ]
# CMD ["/bechamel-api-localdev-server"]


##########
# Configuration for Azure-deployed dev instance
# TODO: For production deployment, build a slimmed-down Docker image
# without development tools as described in above link in "Multi-stage builds"?
FROM base as dev-azuredeploy
# Build the API server with gcflags that disable inlining and optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -o /bechamel-api-server .
# And start the server running
CMD ["/bechamel-api-server"]
