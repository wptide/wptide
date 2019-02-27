FROM alpine:3.8

MAINTAINER XWP <engage@xwp.co>

# Disable Lighthouse error reporting.
ENV CI true

# Copy the shell script.
COPY bin/lh /usr/bin/lh

# Install software.
RUN echo @edge http://nl.alpinelinux.org/alpine/edge/community >> /etc/apk/repositories \
    && apk add --no-cache --upgrade --virtual .build-deps \
        apk-tools \
        nodejs-npm \
    && apk add --no-cache --update --virtual .persistent-deps \
        chromium \
        grep \
        nodejs \
        ttf-freefont \
        udev \
    && npm i -g npm \
    && npm i -g lighthouse \
    && runDeps="$( \
        scanelf --needed --nobanner --recursive /usr/local \
            | awk '{ gsub(/,/, "\nso:", $2); print "so:" $2 }' \
            | sort -u \
            | xargs -r apk info --installed \
            | sort -u \
    )" \
    && apk add --virtual .run-deps $runDeps \
    && apk del .build-deps \
    && rm -rf /var/lib/apk/lists/* /usr/share/doc/* /usr/share/man/* /usr/share/info/* /var/cache/apk/* \
    && chmod +x /usr/bin/lh \
    && mkdir -p /srv/data

# Add the executable.
ADD bin/lighthouse-server /

# Set the working directory.
WORKDIR /srv/data

# Run the executable.
CMD ["/lighthouse-server"]