package main

import (
  "./aegis_redis"
  "fmt"
)

func main() {
  mylist := aegis_redis.Test()

  for i := range mylist {
    fmt.Println(mylist[i])
  }
}
