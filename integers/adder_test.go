package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Adder 2 and 2 should return 4 ", func(t *testing.T) {

		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

	t.Run("Adder 2 and 5 should return 7 ", func(t *testing.T) {

		sum := Add(2, 5)
		expected := 7

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

	t.Run("Adder 5 and 5 should return 10 ", func(t *testing.T) {

		sum := Add(5, 5)
		expected := 10

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
