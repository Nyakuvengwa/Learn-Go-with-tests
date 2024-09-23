package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Anodiwa", "")
		want := "Hello, Anodiwa"

		assertCorrectMessage(got, want, t)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(got, want, t)
	})

	t.Run("In Shona", func(t *testing.T) {
		got := Hello("Anodiwa", "Shona")
		want := "Mhoro Anodiwa"

		assertCorrectMessage(got, want, t)
	})

	t.Run("In Tswana", func(t *testing.T) {
		got := Hello("Anodiwa", "Tswana")
		want := "Dumela Anodiwa"

		assertCorrectMessage(got, want, t)
	})
}

func assertCorrectMessage(got string, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
