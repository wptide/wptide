FROM alpine:3.8

MAINTAINER XWP <engage@xwp.co>

# Install gettext.
RUN apk update \
    && apk add --no-cache \
        gettext \
        bash \
        gawk \
        sed \
        grep \
        bc \
        coreutils \
        ncurses

# Add the executable.
ADD entrypoint.sh /srv/entrypoint.sh

WORKDIR /srv

# Run the executable.
ENTRYPOINT ["./entrypoint.sh"]
