package main

import (
  "net/http"
  "net/url"
  "fmt"

  "github.com/go-martini/martini"
  "github.com/SideCar6/aegis/aegis_redis"
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

  m.Post(api_url + "/keys", func(r *http.Request) (int, string) {
    r.ParseForm()
    body := r.Form["value"][0]
    key := r.Form["key"][0]
    fmt.Println(body, key)

    aegis_redis.SetKey(key, body)
    return 200, "OK"
  })

  m.Run()
}
