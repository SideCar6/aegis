#!/bin/bash

docker rm -f aegis 2&> /dev/null
docker run \
  --name aegis \
  -it --rm \
  -v $HOME/go:/gopath \
  --link go_redis:redis \
  -w /gopath/src/github.com/SideCar6/aegis \
  -h GO \
  -p 3000:3000 \
  jfbrown/golang

