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

url = "http://%s:%s/api/v1/stats" % [ENV['AEGIS_PORT_3000_TCP_ADDR'] || "127.0.0.1", 3000]

res_time = 500
loop do
  path = paths[rand(paths.length)]
  res_time = if res_time.odd?
    res_time - rand(50)
  else
    res_time + rand(50)
  end

  tag = "query-%s" % rand(10)
  timestamp = (("%s.%s" % [ Time.now.to_i, Time.now.strftime("%L")]).to_f * 1000).round

  payload = {
    timestamp: timestamp,
    value: res_time,
    tags: [
      'api_call',
    ],
    meta: {
      query: 'SELECT * FROM somethings ORDER BY something_else DESC',
      name: "API Response Time"
    },
  }.to_json

  query = {
    key: path,
    value: payload,
  }

  puts "--> Inserting key: %s, value: %s" % [ path, payload ]
  HTTParty.post(url, query: query)

  s = rand

  host = `hostname`.chomp

  values = {
    load: (`cat /proc/loadavg`.split(" ")[0].to_f).round,
    "memory:free" => `free | grep Mem | awk '{print $4}'`.chomp,
    "memory:used" => `free | grep Mem | awk '{print $3}'`.chomp,
  }

  values.each_pair do |key, value|
    payload = {
      timestamp: timestamp,
      value: value.to_i,
      tags: [
        'system_info',
      ],
      meta: {
        host: host,
        time: Time.now.to_s,
        name: key,
      },
    }.to_json

    puts "--> Inserting key: %s, value: %s" % [ key, payload ]
    HTTParty.post(url, query: {
      key: "system:%s:%s" % [ key, host ],
      value: payload,
    })
  end

  puts "sleeping for %s seconds..." % s
  sleep(s)
end
