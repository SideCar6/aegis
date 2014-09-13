aegis
=====

Go stats aggregator.

## Usage

To get a redis connection, you must pass in or set `REDIS_PORT_6379_TCP_ADDR`
and `REDIS_PORT_6379_TCP_PORT` environment variables.

    REDIS_PORT_6379_TCP_ADDR=127.0.0.1 REDIS_PORT_6379_TCP_PORT=6379 go run redis-test.go