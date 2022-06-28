package file

import "os"

// F describes a file.
type F struct {
	// Path is the file path.
	Path string
}

// Exists reports whether the file (or directory) exists.
func (f F) Exists() (r bool) {
	if _, err := os.Stat(f.Path); err != nil && os.IsNotExist(err) {
		return
	}
	return true
}

// IsDir reports whether f describes a directory and a
// *PathError if it fails to get f FileInfo. NOTE: on error
// occurrence, the bool value is not reliable since it might
// be a false negative.
func (f F) IsDir() (r bool, err error) {
	info, err := os.Stat(f.Path)
	if err != nil {
		return
	}
	return info.IsDir(), nil
}
