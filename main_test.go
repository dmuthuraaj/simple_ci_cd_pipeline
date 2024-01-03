package main

import "testing"

func TestTellHello(t *testing.T) {
	name := "dhana"
	want := "Hello " + name
	got := TellHello(name)
	if want != got {
		t.Fatalf("Need: %s Got: %s\n", want, got)
	}
}
