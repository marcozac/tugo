package file

import "os"

// F describes a file.
type F struct {
	// Path is the file path.
	Path string
}

// Exists reports whether the file (or directory) exists.
func (f F) Exists() (r bool) {
	if _, err := os.Lstat(f.Path); err != nil && os.IsNotExist(err) {
		return
	}
	return true
}

// IsDir reports whether f describes a directory and a
// *PathError if it fails to get f FileInfo. NOTE: on error
// occurrence, the result is not reliable since it might
// be a false negative.
func (f F) IsDir() (r bool, err error) {
	info, err := os.Lstat(f.Path)
	if err != nil {
		return
	}
	return info.IsDir(), nil
}

// IsSymLink reports whether f describes a symlink and a
// *PathError if it fails to get f FileInfo. NOTE: on error
// occurrence, the result is not reliable since it might
// be a false negative.
func (f F) IsSymLink() (r bool, err error) {
	info, err := os.Lstat(f.Path)
	if err != nil {
		return
	}

	if info.Mode()&os.ModeSymlink != 0 {
		r = true
	}
	return
}
