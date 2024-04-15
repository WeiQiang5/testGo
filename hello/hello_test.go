package hello

import "testing"

func TestHello(t *testing.T) {
	got := returnHello()
	want := "Hello,world"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
