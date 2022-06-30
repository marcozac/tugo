package errc

import (
	"errors"
	"fmt"
	"testing"
)

type TestingError struct{}

func TestErrC(t *testing.T) {
	const msg = "testing new error"
	E := ConstructC[TestingError]()

	t.Run("new_error", func(t *testing.T) {
		err := E.New(msg)
		if err == nil {
			t.Fatal("error not reported")
		}

		if m := err.Error(); m != msg {
			t.Fatalf("expected error message: %q; received: %q\n", msg, m)
		}
	})

	t.Run("new_errorf", func(t *testing.T) {
		err := E.Newf("testing %q with %d specifiers", "message", 2)
		if err == nil {
			t.Fatal("error not reported")
		}

		if want, m := fmt.Sprintf("testing %q with %d specifiers", "message", 2), err.Error(); m != want {
			t.Fatalf("expected error message: %q; received: %q\n", want, m)
		}
	})

	t.Run("is_error", func(t *testing.T) {
		e := E.New(msg)
		if !E.Is(e) {
			t.Fatal("mismatching reported on the same underlying type")
		}
	})

	t.Run("is_not_error", func(t *testing.T) {
		e := errors.New("testing error")
		if E.Is(e) {
			t.Fatal("false positive on underlying error type")
		}
	})
}

func BenchmarkErrC(b *testing.B) {
	const msg = "testing new error"
	E := ConstructC[TestingError]()

	b.Run("new_error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := E.New(msg)
			if err == nil {
				b.Fatal("error not reported")
			}

			if m := err.Error(); m != msg {
				b.Fatalf("expected error message: %q; received: %q\n", msg, m)
			}
		}
	})

	b.Run("is_error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			e := E.New(msg)
			if !E.Is(e) {
				b.Fatal("mismatching reported on the same underlying type")
			}
		}
	})

	b.Run("is_not_error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			e := errors.New("testing error")
			if E.Is(e) {
				b.Fatal("false positive on underlying error type")
			}
		}
	})

}
