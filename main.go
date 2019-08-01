package main
import (
  "flag"
  "fmt"
  "net/http"
  "strings"

  "github.com/go-redis/redis"
)


//package globals
var (
  rc *redis.Client = redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
    Password: "",
    DB: 0,
  })
  genChannel = "genuser"
)

func userProcessor (channelName string, debug bool) {
  var usr *appUser
  psGenerate := rc.Subscribe(channelName)

  for {
    msg, psErr := psGenerate.ReceiveMessage()
    handleError(psErr, "error in subscribe reception for user generation")
    if debug {
      fmt.Printf("received: %s\n", msg.Payload)
    }

    switch msg.Channel {
    case "genuser":
      usr = newAppUser(msg.Payload)
      setErr := rc.Set(usr.Ref, usr.serialize(), 0).Err()
      handleError(setErr, "could not write to key store")
    //case "fetchuser":
    //case "deleteuser":
    }
  }
}

func handleError(e error, msg string) {
  if e != nil {
    //just a stub... there should be more than outputting a message here
    fmt.Println(msg)
  }
}

func usersRoute(w http.ResponseWriter, req *http.Request) {
  username := strings.TrimPrefix(req.URL.Path, "/user/")

  switch req.Method {
  case "DELETE":
    rc.Del("u:" + username)
  case "GET":
    d, _ := rc.Get("u:" + username).Result()
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(d))
  case "POST":
    err := rc.Publish(genChannel, username).Err()
    handleError(err, "could not publish to channel: " + genChannel)
    w.WriteHeader(http.StatusCreated)
  }

  return
}

func main() {
  flag.Parse()
  http.HandleFunc("/user/", usersRoute)

  go userProcessor(genChannel, true)

  fmt.Println("app listening on port 3000")
  if err := http.ListenAndServe("127.0.0.1:3000", nil); err != nil {
    fmt.Println(err.Error())
  }
}
