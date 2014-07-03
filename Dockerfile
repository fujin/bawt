FROM ubuntu
MAINTAINER AJ Christensen <aj@junglist.io>

# Setup Go 1.3
ADD http://golang.org/dl/go1.3.linux-amd64.tar.gz /tmp/go.tgz
RUN tar -C /usr/local -xzf /tmp/go.tgz

# Persistency
VOLUME /bawt
# Wat
ADD . /bawt/go/src/github.com/fujin/bawt
WORKDIR /bawt/go/src/github.com/fujin/bawt

# Populate /bawt/go/bawt
ENV GOROOT /usr/local/go
ENV GOPATH /bawt/go
ENV PATH /usr/local/go/bin:$PATH

RUN apt-get update -qq
RUN apt-get install git bzr mercurial -yqq
RUN go get -v

ENTRYPOINT ["/bawt/go/bin/bawt"]
