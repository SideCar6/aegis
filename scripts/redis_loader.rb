require 'httparty'
require 'json'


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

  payload = {
    timestamp: Time.now.to_i,
    elapsed: elapsed,
    tags: [
      'api_call',
    ],
    meta: {
      query: 'SELECT * FROM somethings ORDER BY something_else DESC',
    },
  }.to_json

  query = {
    key: path,
    value: payload,
  }

  HTTParty.get(ENV['GO_PORT_3000_TCP_ADDR'], port: 3000, query: query)

  s = rand

  if s > 0.8
    path = "system:load:#{`hostname`.chomp}"

    payload = {
      timestamp: Time.now.to_i,
      value: `cat /proc/loadavg`.split(" ")[0],
      tags: [
        'system_info',
      ],
      meta: {
        host: `hostname`.chomp,
        time: Time.now.to_s,
      },
    }.to_json

    HTTParty.get(ENV['GO_PORT_3000_TCP_ADDR'], port: 3000, query: {
      key: path,
      value: payload,
    })
  end

  puts "sleeping for %s seconds..." % s
  sleep(s)
end
