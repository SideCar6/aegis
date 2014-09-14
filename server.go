package main

import (
  "log"
  "net/http"
  "strings"
  "net/url"
  "fmt"

  "github.com/googollee/go-socket.io"
  "github.com/SideCar6/aegis/aegis_redis"
)

var chttp = http.NewServeMux()
const api_url string = "/api/v1"

func main() {
  server, err := socketio.NewServer(nil)
  if err != nil {
    log.Fatal(err)
  }
  server.On("connection", func(so socketio.Socket) {
    log.Println("on connection")

    so.Emit("message", "test message")

    so.On("disconnection", func() {
      log.Println("on disconnect")
    })
  })
  server.On("error", func(so socketio.Socket, err error) {
    log.Println("error:", err)
  })

  http.Handle("/socket.io/", server)

  chttp.Handle("/", http.FileServer(http.Dir("./public")))

  http.HandleFunc("/", HomeHandler)

  http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(200)
    w.Write([]byte("Hello"))
  })

  http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(200)
    w.Write([]byte("Test"))
  })

  http.HandleFunc(api_url + "/keys/", func(w http.ResponseWriter, r *http.Request) {
    keys := aegis_redis.GetKeys()
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write(keys.ToJSON())
  })

  http.HandleFunc(api_url + "/stats", func(w http.ResponseWriter, r *http.Request) {
    log.Println(r.Method)
    switch r.Method {
      case "GET": getStats(w, r)
      case "POST": postStats(w, r)
    }
  })

  log.Fatal(http.ListenAndServe(":3000", nil))
}

func getStats(w http.ResponseWriter, r *http.Request) {
  qs := r.URL.Query()
  key, _ := url.QueryUnescape(qs.Get("key"))
  stats := aegis_redis.GetList(key, 0, -1)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
  w.Write(stats.ToJSON())
}

func postStats(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  body := r.Form["value"][0]
  key := r.Form["key"][0]
  fmt.Println(body, key)

  aegis_redis.SetKey(key, body)
  w.WriteHeader(201)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  if strings.Contains(r.URL.Path, ".") {
    chttp.ServeHTTP(w, r)
  } else {
    http.ServeFile(w, r, "./public/index.html")
  }
}
