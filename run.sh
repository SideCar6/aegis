#!/bin/bash

docker run \
  -it --rm \
  -v $HOME/go:/gopath \
  --link go_redis:redis \
  -h GO \
  jfbrown/golang

