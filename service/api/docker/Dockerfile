FROM gcr.io/google-appengine/php72:latest

COPY composer.json ${APP_DIR}/composer.json

RUN /build-scripts/composer.sh

RUN apt-get update -y && apt-get install -y \
    autoconf \
    make \
    gcc

# Install Xdebug
RUN yes | pecl install xdebug \
    && echo "zend_extension=$(find /opt/php${SHORT_VERSION}/lib/x86_64-linux-gnu/extensions/no-debug-non-zts-20170718/ -name xdebug.so)" > /opt/php${SHORT_VERSION}/lib/conf.d/xdebug.ini \
    && echo "xdebug.remote_enable=on" >> /opt/php${SHORT_VERSION}/lib/conf.d/xdebug.ini \
    && echo "xdebug.remote_autostart=off" >> /opt/php${SHORT_VERSION}/lib/conf.d/xdebug.ini

# Install PHPUnit
RUN curl https://phar.phpunit.de/phpunit-6.phar -L -o phpunit.phar \
    && chmod +x phpunit.phar \
    && mv phpunit.phar /usr/local/bin/phpunit