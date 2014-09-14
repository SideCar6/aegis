package main 

import (
  "github.com/go-martini/martini"
  "./aegis_redis"
  "net/http"
  "net/url"
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

  m.Get(api_url + "/keys", func () []byte {
    keys := aegis_redis.GetKeys()
    return keys.ToJSON()
  })

  m.Get(api_url + "/stats", func (r *http.Request) []byte {
    qs := r.URL.Query()
    key, _ := url.QueryUnescape(qs.Get("key"))
    stats := aegis_redis.GetList(key, 0, -1)
    return stats.ToJSON()
  })

  m.Run()
}