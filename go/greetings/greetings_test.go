package greetings

import (
	"regexp"
	"testing"
)

// check correct return value of the Hello
func TestHelloName(t *testing.T) {
	name := "Rizalina"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello(%q) = %q, %v, want match for %q`, name, msg, err, want)
	}
}

// greetings.Hello with an empty string
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
