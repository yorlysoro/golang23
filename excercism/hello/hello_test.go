package main

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	result := "Hello, World!"
	want := regexp.MustCompile(result)
	msg := Hello()
	if !want.MatchString(msg) {
		t.Fatalf(`Hello() = %q, want match for %#q, nil`, msg, want)
	}
}
