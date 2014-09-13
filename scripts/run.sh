#!/bin/bash

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
  'console')
    docker run --rm -it \
      --link go_redis:redis \
      -h REDISCLI \
      dockerfile/redis \
      /bin/bash
    ;;
esac
