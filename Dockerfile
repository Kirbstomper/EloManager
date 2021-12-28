# Use the Go Image to build our application
from golang:1.17.3 as builder

# Copy the working directory to our source directory in Docker
COPY . /src/elomapper

WORKDIR /src/elomapper
# Build application as a static build
# The mount options add a build cache to speed up multiple builds
RUN --mount=type=cache,target=/root/.cache/go-build \
	--mount=type=cache,target=/go/pkg \
	go build -ldflags '-s -w -extldflags "-static"' -tags osusergo,netgo,sqlite_omit_load_extension -o /usr/local/bin/elomapper .

# Download the static build of litestream directly & make executable
# This is done in builder as chmod doubles size
ADD https://github.com/benbjohnson/litestream/releases/download/v0.3.7/litestream-v0.3.7-linux-amd64-static.tar.gz /tmp/litestream.tar.gz
RUN tar -C /usr/local/bin -xzf /tmp/litestream.tar.gz

# Download s6-overlay for process supervision
# Done again in builder for size
ADD https://github.com/just-containers/s6-overlay/releases/download/v2.2.0.3/s6-overlay-amd64-installer /tmp/
RUN chmod +x /tmp/s6-overlay-amd64-installer

# Starts our final image; based on apline for size
FROM alpine


# Copy executable & litestream from builder
COPY --from=builder /usr/local/bin/elomapper /usr/local/bin/elomapper
COPY --from=builder /usr/local/bin/litestream /usr/local/bin/litestream

# Install s6 for process supervision
COPY --from=builder /tmp/s6-overlay-amd64-installer /tmp/s6-overlay-amd64-installer
RUN /tmp/s6-overlay-amd64-installer /

RUN apk add bash


# Create data directory (although this will likely be mounted too)
RUN mkdir -p /data

# Notify Docker that the container wants to expose a port.
EXPOSE 8080

# Copy s6 init & service definitions.
COPY etc/cont-init.d /etc/cont-init.d
COPY etc/services.d /etc/services.d

# Copy Static Site files
COPY views /views
COPY static /static
# Copy Litestream configuration file.
COPY etc/litestream.yml /etc/litestream.yml

# The kill grace time is set to zero because our app handles shutdown through SIGTERM.
ENV S6_KILL_GRACETIME=0

# Sync disks is enabled so that data is properly flushed.
ENV S6_SYNC_DISKS=1

# Run the s6 init process on entry.

ENTRYPOINT [ "/init" ]
