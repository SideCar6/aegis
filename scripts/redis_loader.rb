require 'redis'
require 'json'

redis = Redis.new({
  host: ENV['REDIS_PORT_6379_TCP_ADDR'],
  port: ENV['REDIS_PORT_6389_TCP_PORT'],
})

loop do
  redis.rpush('GET:/api/v1/users/:id', {
    timestamp: Time.now.to_i,
    elapsed: rand(100),
    tag: "query-%s" % rand(10),
    meta: {},
  }.to_json)

  s = rand(0)
  puts "sleeping for %s seconds..." % s
  sleep(s)
end

