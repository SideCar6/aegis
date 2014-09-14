package main 

import (
  "github.com/go-martini/martini"
  "./aegis_redis"
  // "encoding/json"
)

const api_url string = "/api/v1"

func main() {
  m := martini.Classic()

  m.Get("/hello", func () string {
    return "Hello"
  })
  m.Get("/test", func () string {
    return "Test"
  })
  // m.Get(api_url + "/speed", func () []byte {
  //   j, _ := json.Marshal(aegis_redis.Test())
  //   return j
  // })

  m.Get(api_url + "/keys", func () []byte {
    keys := aegis_redis.GetKeys()
    return keys.ToJSON()
  })

  m.Get(api_url + "/stats/:key", func (params martini.Params) []byte {
    stats := aegis_redis.GetList(params["key"], 0, -1)
    return stats.ToJSON()
  })

  m.Run()
}