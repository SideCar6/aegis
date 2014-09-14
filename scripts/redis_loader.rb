require 'redis'
require 'json'

redis = Redis.new({
  host: ENV['REDIS_PORT_6379_TCP_ADDR'],
  port: ENV['REDIS_PORT_6389_TCP_PORT'],
})

paths = [
  'GET:/api/v1/users/:id',
  'POST:/api/v1/users',
  'PUT:/api/v1/users/:id',
  'DELETE:/api/v1/users/:id',
  'PATCH:/api/v1/users/:id',
  'GET:/api/v1/users/:user_id/cars/:id',
  'GET:/api/v2/cars',
  'GET:/',
  'POST:/accounts/:account_id/users/:id',
]

loop do
  path = paths[rand(paths.length)]
  elapsed = rand(1000)
  tag = "query-%s" % rand(10)

  redis.rpush(path, {
    timestamp: Time.now.to_i,
    elapsed: elapsed,
    tags: [
      'api_call',
    ],
    meta: {
      query: 'SELECT * FROM somethings ORDER BY something_else DESC',
    },
  }.to_json)
  redis.ltrim(path, -99, -1)

  s = rand

  if s > 0.8
    redis.rpush("system:load:#{`hostname`.chomp}", {
      timestamp: Time.now.to_i,
      value: `cat /proc/loadavg`.split(" ")[0],
      tags: [
        'system_info',
      ],
      meta: {
        host: `hostname`.chomp,
        time: Time.now.to_s,
      },
    }.to_json)
  end

  puts "sleeping for %s seconds..." % s
  sleep(s)
end

