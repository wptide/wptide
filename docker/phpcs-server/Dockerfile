FROM alpine:3.7

MAINTAINER XWP <engage@xwp.co>

ENV VENDOR /root/.composer/vendor

RUN set -x \
    && apk add --no-cache \
        php7 \
        php7-ctype \
        php7-simplexml \
        php7-tokenizer \
        php7-xmlwriter \
        ca-certificates \
    && apk add --no-cache -t .build-deps \
        curl \
        php7-json \
        php7-openssl \
        php7-phar \
    && curl -Lo /usr/local/bin/composer https://getcomposer.org/composer.phar \
    && chmod +x /usr/local/bin/composer \
    && composer global require "squizlabs/php_codesniffer=*" \
    && composer global require "wp-coding-standards/wpcs=*" \
    && composer global require "wimg/php-compatibility=*" \
    && $VENDOR/bin/phpcs --config-set show_progress 1 \
    && $VENDOR/bin/phpcs --config-set colors 1 \
    && $VENDOR/bin/phpcs --config-set installed_paths $VENDOR/wp-coding-standards/wpcs,$VENDOR/wimg/php-compatibility \
    && rm /usr/local/bin/composer \
    && apk del --purge .build-deps \
    && mkdir -p /phpcs

ENV PATH=$PATH:$VENDOR/bin/

# Add the executable.
ADD bin/phpcs-server /

# Run the executable.
CMD ["/phpcs-server"]