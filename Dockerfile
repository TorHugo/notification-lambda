# Dockerfile for Go Application using Multi-stage Build
#
# This Dockerfile is designed to build and run a Go application using a multi-stage build process.
# The multi-stage build helps in creating a smaller and more secure final Docker image by separating
# the build environment from the runtime environment.
#
# Stages:
# 1. Builder Stage;
# 2. Final Stage;

# Benefits:
#
# Reduced Image Size: The final image is minimal since it only includes the binary and necessary runtime libraries.
# Improved Security: By excluding build tools and other unnecessary files, the attack surface of the image is reduced.
# Faster Deployment: Smaller image sizes lead to quicker pull times and faster startup, improving overall deployment efficiency.

FROM golang:1.20-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]

# docker build -t [username]/[image_name]:[version] .
# docker push [username]/[image_name]:[version]