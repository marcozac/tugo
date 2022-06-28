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

func BenchmarkCheck(b *testing.B) {
	tmp := b.TempDir()
	c := Check(tmp)

	b.Run("exists", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bExists(b, c)
		}
	})

	b.Run("is_dir", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bIsDir(b, c)
		}
	})

	b.Run("all", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bExists(b, c)
			bIsDir(b, c)
		}
	})
}

func bExists(b *testing.B, c ICheck) {
	if !c.Exists() {
		b.Fatal("false on existing file")
	}
}

func bIsDir(b *testing.B, c ICheck) {
	r, err := c.IsDir()
	if err != nil {
		b.Fatal(err)
	}

	if !r {
		b.Fatal("false on directory")
	}
}
