FROM scratch

MAINTAINER XWP <engage@xwp.co>

# Add CA Certs.
ADD docker/sync-server/cacert.pem /etc/ssl/certs/

# Add the executable.
ADD bin/sync-server /

# Run the executable.
CMD ["/sync-server"]