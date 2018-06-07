FROM gcr.io/google-appengine/php72:latest

COPY composer.json ${APP_DIR}/composer.json

RUN /build-scripts/composer.sh

RUN apt-get update && apt-get install -y autoconf make gcc

# Get and make mongodb PHP driver
RUN yes | apt-get install libssl-dev \
	&& git clone https://github.com/mongodb/mongo-php-driver.git --recursive \
	&& cd mongo-php-driver \
	&& phpize \
	&& ./configure --with-mongodb-ssl=openssl \
	&& make all \
	&& make install \
	&& echo "extension=mongodb.so" > /opt/php${SHORT_VERSION}/lib/conf.d/ext-mongodb.ini \
	&& cd .. \
	&& rm -rf mongo-php-driver

# Install Xdebug
RUN yes | pecl install xdebug \
    && echo "zend_extension=$(find /opt/php${SHORT_VERSION}/lib/x86_64-linux-gnu/extensions/no-debug-non-zts-20170718/ -name xdebug.so)" > /opt/php${SHORT_VERSION}/lib/conf.d/xdebug.ini \
    && echo "xdebug.remote_enable=on" >> /opt/php${SHORT_VERSION}/lib/conf.d/xdebug.ini \
    && echo "xdebug.remote_autostart=off" >> /opt/php${SHORT_VERSION}/lib/conf.d/xdebug.ini

# Install PHPUnit
RUN curl https://phar.phpunit.de/phpunit-6.phar -L -o phpunit.phar \
    && chmod +x phpunit.phar \
    && mv phpunit.phar /usr/local/bin/phpunit