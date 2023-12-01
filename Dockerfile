# syntax=docker/dockerfile:1

##
## Build the application from source
##

FROM golang:1.19 AS build-stage

WORKDIR /app/

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN pwd
RUN ls -al
RUN CGO_ENABLED=0 GOOS=linux go build -o /urs_bin ./internal/cmd/main.go
RUN ls -al

##
## Run the tests in the container
##

FROM build-stage AS run-test-stage
RUN pwd
RUN ls -al
RUN go test -v ./...

##
## Deploy the application binary into a lean image
##

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

# RUN pwd
# RUN ls -al
COPY --from=build-stage /urs_bin /urs_bin
COPY --from=build-stage /app/.env /.env

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/urs_bin"]