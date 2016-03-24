FROM golang:onbuild
MAINTAINER Masashi Shibata<contact@c-bata.link>

ADD . /go/src/github.com/c-bata/gosearch
RUN go install github.com/c-bata/gosearch

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/gosearch

EXPOSE 8080
