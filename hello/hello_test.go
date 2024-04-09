// hello_test.go
package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("amir")
	want := "Mello, amir!"

	if got != want {
		t.Errorf("Wanted %q but got %q", want, got)
	}
}
