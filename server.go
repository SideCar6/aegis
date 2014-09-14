package main

import (
  "log"
  "net/http"
  "strings"
  "net/url"
  "fmt"
  "github.com/SideCar6/aegis/aegis_redis"
  "code.google.com/p/go.net/websocket"
)

var chttp = http.NewServeMux()
const api_url string = "/api/v1"

var openSockets []*websocket.Conn
func main() {
  http.Handle("/websockets", websocket.Handler(socketServer))

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
      case "GET":
        getStats(w, r)
      case "POST":
        postStats(w, r)

    }
  })

  log.Fatal(http.ListenAndServe(":3000", nil))
}

func socketServer(ws *websocket.Conn) {
  fmt.Println("New connection")
  openSockets = append(openSockets, ws)
  buf := make([]byte, 100)
  for {
    mess, err := ws.Read(buf)
    if err == nil {
      fmt.Println(string(buf[0:mess]))
      continue
    }

    for i, s := range openSockets {
      if ws != s {
        continue
      } else {
        openSockets = append(openSockets[:i], openSockets[i + 1:]...)
      }
    }
    fmt.Println("%s", err.Error())
    break
  }
  fmt.Println("Closing connection")
}

func sendMessage(msg string) {
  for _, ws := range openSockets {
    ws.Write([]byte(msg))
  }
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
  sendMessage("{\"" + key + "\":" + body + "}")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  if strings.Contains(r.URL.Path, ".") {
    chttp.ServeHTTP(w, r)
  } else {
    http.ServeFile(w, r, "./public/index.html")
  }
}
