FROM golang:alpine

ENV PATH /go/bin:$PATH

RUN apk add --no-cache --virtual .build-deps \
        glide \
        git

RUN go get github.com/xwp/go-tide/cmd/...
RUN cd /go/src/github.com/xwp/go-tide && glide install
RUN go install -v -x github.com/xwp/go-tide/cmd/...
