FROM alpine:3.8

MAINTAINER XWP <engage@xwp.co>

ENV VENDOR                   /root/.composer/vendor
ENV PHPCS                    3.4.2
ENV WPCS                     2.1.1
ENV PHPCOMPAT                9.1.1
ENV PHPCOMPAT_WP             2.0.0
ENV PHPCOMPAT_PASSWORDCOMPAT 1.0.0
ENV PHPCOMPAT_PARAGONIE      1.0.1
ENV PHPCOMPAT_SYMFONY        1.0.0

RUN set -x \
    && apk add --no-cache \
        php7 \
        php7-ctype \
        php7-json \
        php7-simplexml \
        php7-tokenizer \
        php7-xmlwriter \
        ca-certificates \
    && apk add --no-cache -t .build-deps \
        curl \
        git \
        php7-openssl \
        php7-phar \
    && curl -Lo /usr/local/bin/composer https://getcomposer.org/composer.phar \
    && chmod +x /usr/local/bin/composer \
    && composer global require squizlabs/php_codesniffer:$PHPCS \
    && composer global require wp-coding-standards/wpcs:$WPCS \
    && composer global require phpcompatibility/php-compatibility:$PHPCOMPAT \
    && composer global require phpcompatibility/phpcompatibility-wp:$PHPCOMPAT_WP \
    && composer global require phpcompatibility/phpcompatibility-passwordcompat:$PHPCOMPAT_PASSWORDCOMPAT \
    && composer global require phpcompatibility/phpcompatibility-paragonie:$PHPCOMPAT_PARAGONIE \
    && composer global require phpcompatibility/phpcompatibility-symfony:$PHPCOMPAT_SYMFONY \
    && $VENDOR/bin/phpcs --config-set show_progress 1 \
    && $VENDOR/bin/phpcs --config-set colors 1 \
    && $VENDOR/bin/phpcs --config-set installed_paths $VENDOR/wp-coding-standards/wpcs,$VENDOR/phpcompatibility/php-compatibility,$VENDOR/phpcompatibility/phpcompatibility-wp,$VENDOR/phpcompatibility/phpcompatibility-passwordcompat,$VENDOR/phpcompatibility/phpcompatibility-paragonie,$VENDOR/phpcompatibility/phpcompatibility-symfony \
    && rm /usr/local/bin/composer \
    && apk del --purge .build-deps \
    && mkdir -p /srv/data \
    && mkdir -p /usr/local/etc/php/conf.d \
    && echo 'memory_limit = -1' >> /etc/php7/conf.d/memory-limit.ini

ENV PATH=$PATH:$VENDOR/bin/

# Add the executable.
ADD bin/phpcs-server /

# Set the working directory.
WORKDIR /srv/data

# Run the executable.
CMD ["/phpcs-server"]