package errc

import "fmt"

// ConstructC is a constructor for custom errors with a
// defined underlying type. It returns an ErrC interface, that
// wraps the same methods of Err (as in Construct), but with
// an underlying type T and an additional Is method that reports
// whether an error has T as underlying type.
func ConstructC[T any]() ErrC {
	return errc[T]{}
}

// ErrC is the interface that wraps a custom error methods with a specific
// underlying type.
type ErrC interface {
	Err

	// Is reports whether its underlying type is the same of
	// the e one.
	Is(e error) bool
}

type errc[T any] struct {
	msg string
}

func (errc[T]) New(s string) error {
	return &errc[T]{
		msg: s,
	}
}

func (errc[T]) Newf(s string, m ...any) error {
	return &errc[T]{
		msg: fmt.Sprintf(s, m...),
	}
}

func (errc[T]) Is(e error) (v bool) {
	switch e.(type) {
	case *errc[T]:
		v = true
	}
	return
}

func (e *errc[T]) Error() string {
	return e.msg
}
