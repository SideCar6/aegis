package main 

import (
  "github.com/go-martini/martini"
  "math/rand"
  "strconv"
)

func main() {
  m := martini.Classic()
  r := rand.New(rand.NewSource(99))

  m.Get("/hello", func () string {
    return "Hello"
  })
  m.Get("/test", func () string {
    return "Test"
  })
  m.Get("/api/v1/speed", func () string {
    return strconv.Itoa(r.Intn(100))
  })
  m.Run()
}