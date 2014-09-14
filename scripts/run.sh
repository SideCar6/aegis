#!/bin/bash

if [ `uname | grep Darwin | wc -l` -gt 0 ]; then
  echo '--> Make sure redis is running...'
  bundle --path vendor/bundle
  REDIS_PORT_6379_TCP_ADDR=127.0.01 REDIS_PORT_6379_TCP_PORT=6379 bundle exec ruby redis_loader.rb
fi

if [ $# -lt 1 ]; then
  echo "Usage: $0 <redis|loader|console>"
  exit
fi

case "$1" in
  'loader')
    docker run --rm -it \
      --link aegis_redis:redis \
      --link aegis:aegis \
      -v $(pwd):/data \
      -h redisloader \
      dockerfile/ruby \
      bundle exec /usr/bin/ruby redis_loader.rb
    ;;
  'shell')
    docker run --rm -it \
      --link aegis_redis:redis \
      --link aegis:aegis \
      -v $(pwd):/data \
      -h AEGISRUBY \
      dockerfile/ruby \
      /bin/bash
    ;;
esac
