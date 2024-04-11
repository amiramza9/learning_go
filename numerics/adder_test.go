package numerics

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(9, 7)
	expected := 16

	if got != expected {
		t.Errorf("expected '%d' but got '%d'", expected, got)
	}

}
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)

}
