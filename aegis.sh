#!/bin/bash

if [ `uname | grep Darwin | wc -l` -gt 0 ]; then
  echo "!-> You can't do this on OS X!"
fi

if [ $# -lt 1 ]; then
  echo "Usage: $0 <redis|loader|console>"
  exit
fi

case "$1" in
  'stop')
    docker rm -f aegis 2&> /dev/null
    ;;
  'start')
    docker run -d \
      --name aegis_redis \
      dockerfile/redis

    docker run -it --rm \
      --name aegis \
      --link aegis_redis:redis \
      -p 3000:3000 \
      sidecar6/aegis
    ;;
  'restart')
    $0 stop
    $0 start
    ;;
  'redis-shell')
    docker run --rm -it \
      --link aegis_redis:redis \
      --link aegis:aegis \
      -h REDISCLI \
      dockerfile/redis \
      /bin/bash
    ;;
esac
