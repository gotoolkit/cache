package lru

import "testing"

func TestGet(t *testing.T) {
	lru := New(0)
	lru.Add("myKey", 1234)
	val, ok := lru.Get("myKey1")
	if ok != true {
		t.Fatalf("%s: cache hit =%v; want %v", "hit", ok, !ok)
	} else if ok && val != 1234 {
		t.Fatalf("%s expected get to return 1234 but got %v", "hit", val)
	}
}
