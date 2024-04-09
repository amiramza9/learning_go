// hello_test.go
package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello ", func(t *testing.T) {
		got := Hello("amir")
		want := "Mello, amir!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying Hello World when no argument is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Mello, World!"
		assertCorrectMessage(t, got, want)
	})
}
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Wanted %q but got %q", want, got)
	}
}
