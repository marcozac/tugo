package file

import (
	"os"
	"testing"
)

func TestCheckExists(t *testing.T) {
	tmp := t.TempDir()
	f, err := os.CreateTemp(tmp, "testing_file")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	t.Run("directory_exists", func(t *testing.T) {
		if !Exists(tmp) {
			t.Fatal("false on existing directory")
		}
	})

	t.Run("file_exists", func(t *testing.T) {
		if !Exists(f.Name()) {
			t.Fatal("false on existing file")
		}
	})

	t.Run("file_not_exists", func(t *testing.T) {
		if Exists("fake_testing_file") {
			t.Fatal("true on not existing file")
		}
	})

	t.Run("file_exists_other_error", func(t *testing.T) {
		var Exists = func(f string) bool {
			err = os.ErrInvalid // force different error
			if err != nil {
				if os.IsNotExist(err) {
					return false
				}
			}
			return true
		}
		if !Exists(f.Name()) {
			t.Fatal("false on existing file")
		}
	})
}

func TestCheckIsDir(t *testing.T) {
	tmp := t.TempDir()

	t.Run("is_dir", func(t *testing.T) {
		r, err := IsDir(tmp)
		if err != nil {
			t.Fatal(err)
		}

		if !r {
			t.Fatal("false on directory")
		}
	})

	t.Run("is_not_dir", func(t *testing.T) {
		f, err := os.CreateTemp(tmp, "testing_file")
		if err != nil {
			t.Fatal(err)
		}
		f.Close()

		r, err := IsDir(f.Name())
		if err != nil {
			t.Fatal(err)
		}

		if r {
			t.Fatal("true on non directory")
		}
	})

	t.Run("not_existing_err", func(t *testing.T) {
		_, err := IsDir("fake_dir")
		if err == nil {
			t.Fatal("not existing file error not reported")
		}
	})
}

func TestCheckIsSymlink(t *testing.T) {
	tmp := t.TempDir()
	f, err := os.CreateTemp(tmp, "testing_file")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	sln := f.Name() + "_link"
	if err := os.Symlink(f.Name(), sln); err != nil {
		t.Fatal(err)
	}

	t.Run("is_symlink", func(t *testing.T) {
		r, err := IsSymLink(sln)
		if err != nil {
			t.Error(err)
		}

		if !r {
			t.Error("false on symlink")
		}
	})

	t.Run("is_not_symlink", func(t *testing.T) {
		r, err := IsSymLink(f.Name())
		if err != nil {
			t.Fatal(err)
		}

		if r {
			t.Fatal("true on non symlink")
		}
	})

	t.Run("not_existing_err", func(t *testing.T) {
		_, err := IsSymLink("fake_symlink")
		if err == nil {
			t.Fatal("not existing file error not reported")
		}
	})
}

func BenchmarkCheck(b *testing.B) {
	tmp := b.TempDir()

	f, err := os.CreateTemp(tmp, "testing_file")
	if err != nil {
		b.Fatal(err)
	}
	f.Close()

	sl := f.Name() + "_link"
	if err := os.Symlink(f.Name(), sl); err != nil {
		b.Fatal(err)
	}
	csl := Check(sl)

	b.Run("exists", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if !csl.Exists() {
				b.Fatal("false on existing file")
			}
		}
	})

	b.Run("is_dir", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r, err := csl.IsDir()
			if err != nil {
				b.Fatal(err)
			}

			if r {
				b.Fatal("true on not directory")
			}
		}
	})

	b.Run("is_symlink", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r, err := csl.IsSymLink()
			if err != nil {
				b.Error(err)
			}

			if !r {
				b.Error("false on symlink")
			}
		}
	})
}
