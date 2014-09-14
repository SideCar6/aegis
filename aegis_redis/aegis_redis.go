package aegis_redis

import (
  "fmt"
  "os"
  "time"
  "encoding/json"

  "github.com/fzzy/radix/redis"
)

type (
  Keys []string
  List []string
)

type Stats struct {
  Timestamp   int               `json:"timestamp"`
  Elapsed     int               `json:"elapsed"`
  Tags        []string          `json:"tags"`
  Meta        map[string]string `json:"meta"`
}

type StatsList []Stats

func (keys Keys) ToJSON() []byte {
  j, err := json.Marshal(keys)
  errHndlr(err)

  return j
}

func (stats StatsList) ToJSON() []byte {
  j, err := json.Marshal(stats)
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

func SetKey(key string, value string) bool {
  rdis := connect()
  defer rdis.Close()

  r := rdis.Cmd("select", 0)
  errHndlr(r.Err)

  r = rdis.Cmd("rpush", key, value)
  errHndlr(r.Err)

  return true
}

func GetList(key string, start int32, stop int32) StatsList {
  rdis := connect()
  defer rdis.Close()

  // select database
  r := rdis.Cmd("select", 0)
  errHndlr(r.Err)

  list, err := rdis.Cmd("lrange", key, start, stop).List()
  errHndlr(err)
  log("GetList", list)

  stats := make(StatsList, len(list), len(list))

  for i, l := range list {
    s :=  Stats{}
    json.Unmarshal([]byte(l), &s)
    stats[i] = s
  }

  return stats
}

func connect() *redis.Client {
  c, err := redis.DialTimeout("tcp", os.Getenv("REDIS_PORT_6379_TCP_ADDR") + ":" + os.Getenv("REDIS_PORT_6379_TCP_PORT"), time.Duration(10)*time.Second)

  errHndlr(err)

  return c
}
