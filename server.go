package main

import (
  "log"
  "net/http"

  "github.com/googollee/go-socket.io"
)

func main() {
  server, err := socketio.NewServer(nil)
  if err != nil {
    log.Fatal(err)
  }
  server.On("connection", func(so socketio.Socket) {
    log.Println("on connection")

    go func() {
      so.Emit("message", "test message")

      return
    }()

    so.On("disconnection", func() {
      log.Println("on disconnect")
    })
  })
  server.On("error", func(so socketio.Socket, err error) {
    log.Println("error:", err)
  })

  http.Handle("/socket.io/", server)
  http.Handle("/", http.FileServer(http.Dir("./public")))


  log.Println("Serving at localhost:3000...")
  log.Fatal(http.ListenAndServe(":3000", nil))
}
