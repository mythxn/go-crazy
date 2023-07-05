package app

import "testing"

func TestHelloWorld(t *testing.T) {
	want := "Hello world!"
	got := helloWorld()
	if got != want {
		t.Errorf("hello_world() = %q, want %q", got, want)
	}
}
