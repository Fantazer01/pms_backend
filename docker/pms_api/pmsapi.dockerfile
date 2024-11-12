# syntax=docker/dockerfile:1
FROM golang:1.22-alpine AS build-stage

WORKDIR /app

COPY ./pms_api /app/pms_api/
COPY ./go.* /app/

#RUN go mod download

RUN GOOS=linux go build -o entrypoint_pms_api pms_api/cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=build-stage /app/entrypoint_pms_api /app/entrypoint_pms_api
COPY --from=build-stage /app/pms_api/config/config.yaml /app/pms_api/config/config.yaml

# Run
CMD ["/app/entrypoint_pms_api"]
