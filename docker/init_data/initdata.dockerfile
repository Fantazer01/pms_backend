# syntax=docker/dockerfile:1
FROM golang:1.22-alpine AS build-stage

WORKDIR /app

COPY ./init_data /app/init_data/
COPY ./go.* /app/

#RUN go mod download

RUN GOOS=linux go build -o entrypoint_init_data init_data/main.go

FROM scratch
WORKDIR /app
COPY --from=build-stage /app/entrypoint_init_data /app/entrypoint_init_data

# Run
CMD ["/app/entrypoint_init_data"]
