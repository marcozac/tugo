package file

// ICheck is the interface that wraps F check methods.
type ICheck interface {
	// Exists reports whether the file (or directory) exists.
	Exists() bool

	// IsDir reports whether f describes a directory and a
	// *PathError if it fails to get f FileInfo. NOTE: on error
	// occurrence, the bool value is not reliable since it might
	// be a false negative.
	IsDir() (bool, error)
}

// Check returns a ICheck interface with underlying implementation
// of F.Path: f.
func Check(f string) ICheck {
	return F{
		Path: f,
	}
}

// Exists is a shortcut for ICheck.Exists().
func Exists(f string) (r bool) {
	return Check(f).Exists()
}

// IsDir is a shortcut for ICheck.IsDir().
func IsDir(f string) (r bool, err error) {
	return Check(f).IsDir()
}
