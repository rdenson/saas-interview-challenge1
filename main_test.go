package main
import (
  "math/rand"
  "testing"
  "time"

  "github.com/go-redis/redis"
)


func BenchmarkUserProcessor(b *testing.B) {
  broker := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
    Password: "",
    DB: 0,
  })
  testChannel := "benchmark-test"
  randomSource := rand.NewSource(time.Now().UnixNano())
  nri := rand.New(randomSource)

  go userProcessor(testChannel, false)
  for i:=0; i<b.N; i++ {
    newAppUser(string(nri.Intn(25) + 65))
    broker.Publish(testChannel, string(nri.Intn(25) + 65))
  }
}
