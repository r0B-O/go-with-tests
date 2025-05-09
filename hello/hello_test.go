package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("r0B-O")
	want := "Hello, r0B-O"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
