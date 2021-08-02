package skiplist

import (
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	New(18)
}

func TestCorrectness(t *testing.T) {
	s := New(18)
	batch := 100 * 10000
	for i := 0; i < batch; i++ {
		s.Set(i, i)
	}
	for i := 0; i < batch; i++ {
		v, ok := s.Get(i)
		if !ok || v != i {
			t.Fatal("get value not equal to ", i)
		}
	}
	for i := 0; i < batch/2; i++ {
		if !s.Del(i) {
			t.Fatal("del err at", i)
		}
	}
	for i := 0; i < batch/2; i++ {
		_, ok := s.Get(i)
		if ok {
			t.Fatal("get deleted value err", i)
		}
	}
	for i := batch / 2; i < batch; i++ {
		v, ok := s.Get(i)
		if !ok || v != i {
			t.Fatal("get not deleted value err", i)
		}
	}
}

func TestParallel(t *testing.T) {
	s := New(18)
	batch := 100 * 10000

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < batch/2; i++ {
			s.Set(i, i)
		}
		wg.Done()
	}()
	go func() {
		for i := batch / 4; i < batch; i++ {
			s.Set(i, i)
		}
		wg.Done()
	}()
	wg.Wait()
	for i := 0; i < batch; i++ {
		v, ok := s.Get(i)
		if !ok || v != i {
			t.Fatal("get value not equal to ", i)
		}
	}
	wg.Add(2)
	go func() {
		for i := 0; i < batch/2; i++ {
			s.Del(i)
		}
		wg.Done()
	}()
	go func() {
		for i := batch / 4; i < batch; i++ {
			s.Del(i)
		}
		wg.Done()
	}()
	wg.Wait()
	for i := 0; i < batch; i++ {
		_, ok := s.Get(i)
		if ok {
			t.Fatal("get deleted value err", i)
		}
	}
	wg.Wait()
}
