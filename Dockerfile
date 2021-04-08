FROM golang:1.15-buster

WORKDIR /build


# Build our tweaked nomad plugin for the time being
COPY nomad /build/nomad
# COPY nomad-test /build/nomad-test
# COPY config /build/config
# COPY scripts /build/scripts
# COPY docs /build/docs
COPY go.* /build
COPY *.go /build
RUN go build -trimpath -ldflags="-w -s" -o /build/steampipe-plugin-nomad.plugin  *.go
