package errc

import "fmt"

// Construct is a constructor for custom errors.
// It returns Err interface, that wraps New and Newf methods to
// generate errors from a given text.
func Construct() Err {
	return err{}
}

// Err is the interface that wraps a custom error methods.
type Err interface {
	// New instances a new error with s text. As its homonym
	// in errors package, each call generates a different error,
	// even using the same text.
	New(s string) error

	// Newf instances a new error as New formatting the text with Sprintf.
	Newf(s string, m ...any) error
}

// err implements custom error methods.
type err struct {
	// msg is the error text.
	msg string
}

func (err) New(s string) error {
	return &err{
		msg: s,
	}
}

func (err) Newf(s string, m ...any) error {
	return &err{
		msg: fmt.Sprintf(s, m...),
	}
}

func (e *err) Error() string {
	return e.msg
}
