package skiplist

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

type SkipList struct {
	limit  int
	level  int
	size   int
	header *entry
	rand   *rand.Rand
}

type entry struct {
	k     interface{}
	p     unsafe.Pointer
	hash  uint64
	level int
	next  []*entry
}

func newEntry(k, v interface{}, level int) *entry {
	return &entry{k: k, p: unsafe.Pointer(&v), hash: hash(k), level: level, next: make([]*entry, level+1)}
}

func (e *entry) value() interface{} {
	return *(*interface{})(e.p)
}

func New(limit int) SkipList {
	return SkipList{
		limit:  limit,
		header: newEntry(nil, nil, limit),
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s SkipList) random() int {
	i := 0
	for i < s.limit && s.rand.Float64() <= 0.25 {
		i++
	}
	return i
}

func (s *SkipList) Set(k, v interface{}) {
	e := newEntry(k, v, s.random())
	if s.size == 0 {
		s.level = e.level
		s.size++
		s.header.next[0] = e
		return
	}
	prev := s.header
	pres := make([]*entry, e.level+1)
	for l := max(s.level, e.level); l >= 0; l-- {
		next := prev.next[l]
		for next != nil && e.hash >= next.hash {
			if e.k == next.k {
				next.p = e.p
				return
			}
			prev = next
			next = next.next[l]
		}
		if l <= e.level {
			pres[l] = prev
		}
	}
	for i := 0; i < len(pres); i++ {
		e.next[i] = pres[i].next[i]
		pres[i].next[i] = e
	}
	s.size++
	s.level = max(s.level, e.level)
}

func (s SkipList) Get(k interface{}) (interface{}, bool) {
	prev := s.header
	h := hash(k)
	for l := s.level; l >= 0; l-- {
		next := prev
		for next = next.next[l]; next != nil && h >= next.hash; {
			if k == next.k {
				return next.value(), true
			}
			prev = next
			next = next.next[l]
		}
	}
	return nil, false
}

func (s SkipList) print() string {
	buf := bytes.Buffer{}
	for i := s.level; i >= 0; i-- {
		buf.WriteString("||->")
		next := s.header.next[i]
		for next != nil {
			buf.WriteString(fmt.Sprintf("%v", next.value()))
			buf.WriteString("->")
			next = next.next[i]
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
