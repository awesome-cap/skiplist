package skiplist

import (
	"github.com/abbychau/jumplist"
	"testing"
	"time"
)

func TestSkipList_Hash(t *testing.T) {
	batch := 1000000
	start := time.Now().UnixNano()
	for i := 0; i < batch; i++ {
		hash(i)
	}
	end := time.Now().UnixNano()
	t.Log("set time: ", (end-start)/1e6)
}

func TestSkipList_Random(t *testing.T) {
	s := New(18)
	batch := 1000000
	start := time.Now().UnixNano()
	for i := 0; i < batch; i++ {
		s.random()
	}
	end := time.Now().UnixNano()
	t.Log("set time: ", (end-start)/1e6)
}

func TestSkipList_Set2(t *testing.T) {
	s := New(18)
	batch := 100
	start := time.Now().UnixNano()
	for i := 0; i < batch; i++ {
		s.Set(i, i)
	}
	end := time.Now().UnixNano()
	t.Log("set time: ", (end-start)/1e6)

	start = time.Now().UnixNano()
	for i := 0; i < batch; i++ {
		v, ok := s.Get(i)
		if !ok {
			t.Fatal("get err", i)
		}
		if v != i {
			t.Fatal("v not equal i")
		}
	}
	end = time.Now().UnixNano()
	t.Log("get time: ", (end-start)/1e6)
}

func TestSkipList_Set(t *testing.T) {
	s := New(18)
	batch := 1000000
	start := time.Now().UnixNano()
	for i := 0; i < batch; i++ {
		s.Set(i, i)
	}
	end := time.Now().UnixNano()
	t.Log("set time: ", (end-start)/1e6)

	start = time.Now().UnixNano()
	for i := 0; i < batch; i++ {
		v, ok := s.Get(i)
		if !ok {
			t.Fatal("get err", i)
		}
		if v != i {
			t.Fatal("v not equal i")
		}
	}
	end = time.Now().UnixNano()
	t.Log("get time: ", (end-start)/1e6)
}

func TestJumpList_Set(t *testing.T) {
	s := jumplist.New()
	batch := 1000000
	start := time.Now().UnixNano()
	for i := 0; i < batch; i++ {
		s.Set(float64(i), i)
	}
	end := time.Now().UnixNano()
	t.Log("set time: ", (end-start)/1e6)

	start = time.Now().UnixNano()
	for i := 0; i < batch; i++ {
		s.Get(float64(i))
	}
	end = time.Now().UnixNano()
	t.Log("get time: ", (end-start)/1e6)
}

func BenchmarkSkipList_Set(b *testing.B) {
	b.ResetTimer()
	s := New(48)
	batch := 1000000
	for i := 0; i < batch; i++ {
		s.Set(i, i)
	}
}
