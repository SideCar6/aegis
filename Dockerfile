FROM        google/golang
MAINTAINER  SideCar6, Jack Brown & Matt Merkes

RUN         mkdir -p /gopath/src/aegis
ADD         . /gopath/src/aegis/.
WORKDIR     /gopath/src/aegis
RUN         go get

EXPOSE      3000

ENTRYPOINT  ["go"]
CMD         ["run", "server.go"]
