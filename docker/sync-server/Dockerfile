FROM golang:alpine

RUN apk add --no-cache --virtual .build-deps \
        glide \
        git

RUN git clone https://github.com/xwp/go-tide.git /go/src/github.com/xwp/go-tide

RUN cd /go/src/github.com/xwp/go-tide && glide install