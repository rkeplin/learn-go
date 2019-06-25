package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Expected %d, but got %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAddMulti(t *testing.T) {
	got := AddMulti(1, 2, 3, 4, 5)
	want := 15

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
