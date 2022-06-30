package errc

import (
	"fmt"
	"testing"
)

func TestErr(t *testing.T) {
	const msg = "testing new error"
	E := Construct()

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
}

func BenchmarkErr(b *testing.B) {
	const msg = "testing new error"
	for i := 0; i < b.N; i++ {
		E := Construct()

		err := E.New(msg)
		if err == nil {
			b.Fatal("error not reported")
		}

		if m := err.Error(); m != msg {
			b.Fatalf("expected error message: %q; received: %q\n", msg, m)
		}
	}
}
