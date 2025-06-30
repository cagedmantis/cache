package lru_test

import (
	"testing"

	"argc.dev/cache/lru"
)

func TestGet(t *testing.T) {
	c := lru.New[string, string](1)
	key1, key2, want1, want2 := "one", "two", "one-val", "two-val"
	c.Put(key1, want1)
	c.Put(key2, want2)
	if got, ok := c.Get(key2); got != want2 || !ok {
		t.Errorf("Get(%s) = %q, %t; want %q, true", key2, got, ok, want2)
	}
	if got, ok := c.Get(key1); ok {
		t.Errorf("Get(%s) = %q, %t; want nil, false", key1, got, ok)
	}
}

func TestPut(t *testing.T) {
	c := lru.New[int, int](50)
	for idx := 0; idx < 100; idx++ {
		c.Put(idx, idx*10)
	}
	for idx := 0; idx < 100; idx++ {
		got, ok := c.Get(idx)
		if idx < 50 && ok {
			t.Errorf("Put() failed for %+d", idx)
		}
		if idx >= 50 && !ok {
			t.Errorf("Put() failed for %+d=%d", idx, got)
		}
	}
}
