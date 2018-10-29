FROM alpine:3.8

MAINTAINER XWP <engage@xwp.co>

# Add CA Certs & working directory.
RUN apk update \
    && apk add ca-certificates \
    && rm -rf /var/cache/apk/* \
    && mkdir -p /srv/data

# Add the executable.
ADD bin/sync-server /

# Set the working directory.
WORKDIR /srv/data

# Run the executable.
CMD ["/sync-server"]