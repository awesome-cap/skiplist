package skiplist

import (
	"math/rand"
	"sync"
	"time"
)

type SkipList struct {
	sync.Mutex
	limit int
	size  int
	head  *entry
	prev  []*entry
	rand  *rand.Rand
}

type entry struct {
	k    interface{}
	v    interface{}
	hash uint64
	next []*entry
}

func newEntry(k, v interface{}, hash uint64, level int) *entry {
	return &entry{k: k, v: v, hash: hash, next: make([]*entry, level+1)}
}

func New(limit int) *SkipList {
	return &SkipList{
		limit: limit,
		head:  newEntry(nil, nil, 0, limit+1),
		prev:  make([]*entry, limit+1),
		rand:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *SkipList) random() int {
	i := 0
	for i < s.limit && s.rand.Float32() < 0.251 {
		i++
	}
	return i
}

func (s *SkipList) Set(k, v interface{}) {
	s.Lock()
	defer s.Unlock()
	h, prev := hash(k), s.head
	var next *entry
	for l := s.limit; l >= 0; l-- {
		next = prev.next[l]
		for next != nil && h >= next.hash {
			if h == next.hash && next.k == k {
				next.v = v
				return
			}
			prev = next
			next = next.next[l]
		}
		s.prev[l] = prev
	}
	e := newEntry(k, v, h, s.random())
	for i := range e.next {
		e.next[i] = s.prev[i].next[i]
		s.prev[i].next[i] = e
	}
	s.size++
}

func (s *SkipList) Get(k interface{}) (interface{}, bool) {
	prev := s.head
	h := hash(k)
	for l := s.limit; l >= 0; l-- {
		next := prev.next[l]
		for next != nil && h >= next.hash {
			if h == next.hash && next.k == k {
				return next.v, true
			}
			prev = next
			next = next.next[l]
		}
	}
	return nil, false
}

func (s *SkipList) Del(k interface{}) bool {
	s.Lock()
	defer s.Unlock()
	h, prev := hash(k), s.head
	var target, next *entry
	for l := s.limit; l >= 0; l-- {
		next = prev.next[l]
		for next != nil && h >= next.hash {
			if h == next.hash && next.k == k {
				target = next
				break
			}
			prev = next
			next = next.next[l]
		}
		s.prev[l] = prev
	}
	if target == nil {
		return false
	}
	for i := range target.next {
		s.prev[i].next[i] = target.next[i]
	}
	s.size--
	return true
}
