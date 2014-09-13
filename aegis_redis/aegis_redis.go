// example program

package aegis_redis

import (
  "fmt"
  "github.com/fzzy/radix/redis"
  "os"
  "time"
  "encoding/json"
)

type (
  Keys []string
  List []string
)

func (keys Keys) ToJSON() []byte {
  j, err := json.Marshal(keys)
  errHndlr(err)

  return j
}

func (list List) ToJSON() []byte {
  j, err := json.Marshal(list)
  errHndlr(err)

  return j
}

func log(comment string, l []string) {
  for i := range l {
    fmt.Println("[ " + comment + " ]\t" + l[i])
  }
}

func errHndlr(err error) {
  if err != nil {
    fmt.Println("error:", err)
    os.Exit(1)
  }
}

func GetKeys() Keys {
  rdis := connect()
  defer rdis.Close()

  r := rdis.Cmd("select", 0)
  errHndlr(r.Err)

  keys, err := rdis.Cmd("keys", "*").List()
  errHndlr(err)
  log("GetKeys", keys)

  return keys
}

func GetList(key string, start int32, stop int32) List {
  rdis := connect()
  defer rdis.Close()

  // select database
  r := rdis.Cmd("select", 0)
  errHndlr(r.Err)

  list, err := rdis.Cmd("lrange", key, start, stop).List()
  errHndlr(err)
  log("GetList", list)

  return list
}

func connect() *redis.Client {
  c, err := redis.DialTimeout("tcp", os.Getenv("REDIS_PORT_6379_TCP_ADDR") + ":" + os.Getenv("REDIS_PORT_6379_TCP_PORT"), time.Duration(10)*time.Second)

  errHndlr(err)

  return c
}
