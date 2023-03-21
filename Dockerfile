FROM golang:1.20-alpine AS build

WORKDIR /app

# Copy the Go modules (go.mod) and the Go modules checksums (go.sum) files.
COPY ./go.mod ./go.sum ./
# Download modules to local cache.
RUN go mod download

# Copy the source code.
COPY ./ ./
# Build BendoBot.
RUN go build


FROM golang:1.20-alpine AS deploy

# Copy the Go path contents, to have all the dependencies, which are required to run BendoBot.
COPY --from=build /go /go

WORKDIR /app

# Copy the compiled BendoBot.
COPY --from=build /app/BendoBot ./

# Go into a new directory, which should be used as the mount point for a Docker volume.
# This Docker volume should contain a valid configuration file (config.json) for BendoBot.
# This way we can easily use a configuration file inside a Docker container.
WORKDIR ./config_volume

# Run BendoBot.
ENTRYPOINT ["../BendoBot"]
