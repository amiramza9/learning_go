// hello_test.go
package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello ", func(t *testing.T) {
		got := Hello("amir")
		want := "Mello, amir!"

		if got != want {
			t.Errorf("Wanted %q but got %q", want, got)
		}
	})
	t.Run("Saying Hello World when no argument is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Mello, World!"

		if got != want {

			t.Errorf("Wanted %q but got %q", want, got)
		}

	})
}
