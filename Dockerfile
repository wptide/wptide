# @TODO: Build from a Lighthouse image.
FROM scratch
MAINTAINER XWP <engage@xwp.co>

# Add the executable.
ADD bin/lh-server /

# Add certs because we are connecting to some services over SSL/TSL.
ADD certs/cacert.pem /etc/ssl/certs/

# Run the executable.
CMD ["/lh-server"]