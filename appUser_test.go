package main
import (
  "testing"
  "regexp"
)


func TestNewAppUser(t *testing.T) {
  tc1 := newAppUser("foo")
  usernameRe := regexp.MustCompile(`^u:.+`)

  if tc1.CreateDate.IsZero() {
    t.Errorf("expected a \"non-zero\" create date, found:\ncreate date: %s", tc1.CreateDate.String())
  }

  if !usernameRe.MatchString(tc1.Ref) {
    t.Errorf("expected a prefixed user reference, found:\nreference: %s\n", tc1.Ref)
  }

  if !tc1.Active {
    t.Errorf("expected active to be true")
  }
}

func TestGenString(t *testing.T) {
  tc1 := genString()
  pattern := regexp.MustCompile(`([a-z]{4}-){3}([a-z]{4})`)

  if len(tc1) == 16 {
    t.Errorf("expected genString() to return a string of length 16, length found: %d", len(tc1))
  }

  if !pattern.MatchString(tc1) {
    t.Errorf("expected tc1 to match \"([a-z]{4}-){3}([a-z]{4})\", found: %s", tc1)
  }
}

func TestSetRef(t *testing.T) {
  usr := &appUser{}
  expected := "u:foo"

  usr.setRef("foo")
  if usr.Ref != expected {
    t.Errorf("expected the reference to be %s, found: %s", expected, usr.Ref)
  }
}
