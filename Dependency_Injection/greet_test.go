package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// Buffer of bytes
	buffer := bytes.Buffer{}

	// Passes the buffer address and a string
	Greet(&buffer, "Kelvin")

	// The idea of using buffer it's because it implements the writer method
	// So then we can check what was writen
	got := buffer.String()
	want := "Hello, Kelvin"

	if got != want {
		t.Errorf("got %q qant %q", got, want)
	}
}
