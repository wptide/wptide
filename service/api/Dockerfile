FROM gcr.io/google-appengine/php72:latest

# Update repos & Install GCC
RUN apt-get update && apt-get install -y autoconf make gcc

# Install Xdebug
RUN yes | pecl install xdebug \
    && echo "zend_extension=$(find /opt/php72/lib/x86_64-linux-gnu/extensions/no-debug-non-zts-20170718/ -name xdebug.so)" > /opt/php72/lib/conf.d/xdebug.ini \
    && echo "xdebug.remote_enable=on" >> /opt/php72/lib/conf.d/xdebug.ini \
    && echo "xdebug.remote_autostart=off" >> /opt/php72/lib/conf.d/xdebug.ini

# Install PHPUnit
RUN curl https://phar.phpunit.de/phpunit-6.phar -L -o phpunit.phar \
    && chmod +x phpunit.phar \
    && mv phpunit.phar /usr/local/bin/phpunit