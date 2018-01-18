FROM alpine
MAINTAINER XWP <engage@xwp.co>

RUN apk add --no-cache \
	nodejs \
	nodejs-npm \
	grep \
    udev \
    ttf-freefont \
    chromium

# Download latest Lighthouse from npm.
# cache bust so we always get the latest version of LH when building the image.
ARG CACHEBUST=1
RUN npm i lighthouse -g

RUN apk del nodejs-npm


#VOLUME /home/chrome/reports
#WORKDIR /home/chrome/reports

# Disable Lighthouse error reporting to prevent prompt.
ENV CI=true

EXPOSE 8080

# Add the executable.
ADD bin/lh-server /

# Add certs because we are connecting to some services over SSL/TSL.
ADD certs/cacert.pem /etc/ssl/certs/

# Run the executable.
CMD ["/lh-server"]