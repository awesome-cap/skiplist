package skiplist

import (
	"math/rand"
	"sync"
	"time"
	"unsafe"
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
	p    unsafe.Pointer
	hash uint64
	next []*entry
}

func newEntry(k, v interface{}, level int) *entry {
	return &entry{k: k, p: unsafe.Pointer(&v), hash: hash(k), next: make([]*entry, level+1)}
}

func (e *entry) value() interface{} {
	return *(*interface{})(e.p)
}

func (e *entry) level() int {
	return len(e.next) - 1
}

func New(limit int) *SkipList {
	return &SkipList{
		limit: limit,
		head:  newEntry(nil, nil, limit),
		prev:  make([]*entry, limit),
		rand:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *SkipList) random() int {
	i := 0
	for i < s.limit && s.rand.Float64() <= 0.29 {
		i++
	}
	return i
}

func (s *SkipList) Set(k, v interface{}) {
	s.Lock()
	defer s.Unlock()
	h, prev := hash(k), s.head
	for l := s.limit - 1; l >= 0; l-- {
		next := prev.next[l]
		for next != nil && h >= next.hash {
			if h == next.hash && next.k == k {
				next.p = unsafe.Pointer(&v)
				return
			}
			prev = next
			next = next.next[l]
		}
		s.prev[l] = prev
	}
	s.size++
	e := &entry{k: k, p: unsafe.Pointer(&v), hash: h, next: make([]*entry, s.random()+1)}
	for i := 0; i <= e.level(); i++ {
		e.next[i] = s.prev[i].next[i]
		s.prev[i].next[i] = e
	}
}

func (s *SkipList) Get(k interface{}) (interface{}, bool) {
	prev := s.head
	h := hash(k)
	for l := s.limit - 1; l >= 0; l-- {
		next := prev.next[l]
		for next != nil && h >= next.hash {
			if h == next.hash && next.k == k {
				return next.value(), true
			}
			prev = next
			next = next.next[l]
		}
	}
	return nil, false
}
