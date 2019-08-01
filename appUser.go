package main
import (
  "encoding/json"
  "math/rand"
  "strings"
  "time"
)


//unit we're working on
type appUser struct {
  Key string `json:"productKey"`
  Ref string `json:"dbRef"`
  CreateDate time.Time `json:"created"`
  Active bool `json:"active"`
}
func (a *appUser) serialize() string {
  s, marshalErr := json.Marshal(a)
  handleError(marshalErr, "appUser.serialize() failed")

  return string(s)
}
func (a *appUser) setRef(s string) {
  a.Ref = "u:" + s
}

//custom initializer for the work unit
func newAppUser(name string) *appUser {
  usr := &appUser{
    Key: genString(),
    CreateDate: time.Now(),
    Active: true,
  }

  usr.setRef(name)

  return usr
}

//create a custom key identifier
func genString() string {
  const str_sz int = 16
  var newstr strings.Builder
  randomSource := rand.NewSource(time.Now().UnixNano())
  nri := rand.New(randomSource)

  for i:=0; i<str_sz; i++ {
    if (i % 4) == 0 && i != 0 && i != str_sz {
      newstr.WriteString("-")
    }
    newstr.WriteString(string(nri.Intn(25) + 97))
  }

  return newstr.String()
}
