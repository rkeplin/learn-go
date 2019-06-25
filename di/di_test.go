package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Rob")

	got := buffer.String()
	want := "Hello, Rob"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
