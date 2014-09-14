package main

import (
  "fmt"

  "github.com/SideCar6/aegis/aegis_redis"
)

func main() {
  fmt.Println("ALL KEYS: ")
  mykeys := aegis_redis.GetKeys()

  for i := range mykeys {
    fmt.Println(mykeys[i])
    aegis_redis.GetList(mykeys[i], -3, -1)
  }
}
