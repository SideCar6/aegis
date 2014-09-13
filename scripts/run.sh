#!/bin/bash

if [ `uname | grep Darwin | wc -l` -lt 1 ]; then
  echo '--> Make sure redis is running...'
  bundle --path vendor/bundle
  REDIS_PORT_6379_TCP_ADDR=127.0.01 REDIS_PORT_6379_TCP_PORT=6379 bundle exec ruby redis_loader.rb
fi

if [ $# -lt 1 ]; then
  echo "Usage: $0 <redis|loader|console>"
  exit
fi

case "$1" in
  'redis')
    docker rm -f go_redis 2&> /dev/null

    docker run -d \
      --name go_redis \
      dockerfile/redis
    ;;
  'loader')
    docker run --rm -it \
      --link go_redis:redis \
      -v $(pwd):/data \
      dockerfile/ruby \
      bundle exec /usr/bin/ruby redis_loader.rb
    ;;
  'shell')
    docker run --rm -it \
      --link go_redis:redis \
      -v $(pwd):/data \
      dockerfile/ruby \
      /bin/bash
    ;;
  'console')
    docker run --rm -it \
      --link go_redis:redis \
      -h REDISCLI \
      dockerfile/redis \
      /bin/bash
    ;;
esac
