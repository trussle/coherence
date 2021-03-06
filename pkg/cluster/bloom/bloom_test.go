package bloom

import (
	"bytes"
	"testing"
	"testing/quick"

	"github.com/trussle/uuid"
)

func TestBloom(t *testing.T) {
	t.Parallel()

	t.Run("cap", func(t *testing.T) {
		bloom := New(512, 4)
		if expected, actual := uint(512), bloom.Cap(); expected != actual {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
	})

	t.Run("add", func(t *testing.T) {
		fn := func(a uuid.UUID) bool {
			bloom := New(256*2, 4)
			return bloom.Add(a.String()) == nil
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("clear", func(t *testing.T) {
		fn := func(a uuid.UUID) bool {
			bloom := New(256*2, 4)
			if err := bloom.Add(a.String()); err != nil {
				t.Fatal(err)
			}
			return bloom.Clear(a.String()) == nil
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("clear after add", func(t *testing.T) {
		fn := func(a uuid.UUID) bool {
			bloom := New(256*2, 4)
			if err := bloom.Add(a.String()); err != nil {
				t.Fatal(err)
			}
			bloom.Clear(a.String())
			return "{}" == bloom.String()
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("write and read", func(t *testing.T) {
		fn := func(a uuid.UUID) bool {
			bloom := New(256*2, 4)
			if err := bloom.Add(a.String()); err != nil {
				t.Error(err)
			}

			buf := new(bytes.Buffer)
			if _, err := bloom.Write(buf); err != nil {
				t.Error(err)
			}

			other := new(Bloom)
			if _, err := other.Read(buf); err != nil {
				t.Error(err)
			}

			return bloom.String() == other.String()
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("contains nothing", func(t *testing.T) {
		fn := func(a uuid.UUID) bool {
			bloom := New(256*2, 4)
			ok, err := bloom.Contains(a.String())
			if err != nil {
				t.Error(err)
			}
			return !ok
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("contains something", func(t *testing.T) {
		fn := func(a uuid.UUID) bool {
			bloom := New(256*2, 4)
			if err := bloom.Add(a.String()); err != nil {
				t.Fatal(err)
			}

			ok, err := bloom.Contains(a.String())
			if err != nil {
				t.Error(err)
			}
			return ok
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("contains multiple", func(t *testing.T) {
		fn := func(a []uuid.UUID) bool {
			if len(a) == 0 {
				return true
			}

			bloom := New(256*2, 4)
			for _, v := range a {
				if err := bloom.Add(v.String()); err != nil {
					t.Fatal(err)
				}
			}

			for k, v := range a {
				ok, err := bloom.Contains(v.String())
				if err != nil {
					t.Error(err)
				}
				if !ok {
					t.Errorf("expected %q at %d to be inside the bloom", v.String(), k)
					return false
				}
			}
			return true
		}
		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})
}
