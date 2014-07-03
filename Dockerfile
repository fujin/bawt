FROM ubuntu
MAINTAINER AJ Christensen <aj@junglist.io>



# Setup Go 1.3
ADD http://golang.org/dl/go1.3.linux-amd64.tar.gz /tmp/go.tgz
RUN tar -C /usr/local -xzf /tmp/go.tgz

# Persistency
VOLUME /bawt
ADD . /bawt
WORKDIR /bawt

# Populate /bawt/go/bawt
ENV GOPATH /bawt
ENV GOROOT /usr/local/go
ENV PATH /usr/local/go/bin:$PATH

RUN apt-get install git bzr mercurial -yqq
RUN go get -u -v github.com/danryan/hal
RUN go install -v

ENTRYPOINT ["/bawt/bin/bawt"]
