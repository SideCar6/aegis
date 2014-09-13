package main

import (
  "./aegis_redis"
  "fmt"
)

func main() {
  fmt.Println("ALL KEYS: ")
  mykeys := aegis_redis.GetKeys()

  for i := range mykeys {
    fmt.Println(mykeys[i])
    aegis_redis.GetList(mykeys[i], -3, -1)
  }
}
