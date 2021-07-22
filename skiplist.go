package skiplist

import (
	"bytes"
	"fmt"
	"math/rand"
	"unsafe"
)

type SkipList struct {
	limit   int
	headers []*entry
}

type entry struct {
	k    interface{}
	p    unsafe.Pointer
	hash uint64
	next []*entry
}

func (e *entry) level() int {
	return len(e.next) - 1
}

func (e *entry) value() interface{} {
	return *(*interface{})(e.p)
}

func New(limit int) SkipList {
	return SkipList{
		limit:   limit,
		headers: make([]*entry, limit),
	}
}

func (s SkipList) level() int {
	return len(s.headers) - 1
}

func (s SkipList) random() int {
	i := 0
	for i < s.limit && rand.Float32() <= 0.25 {
		i++
	}
	return i
}

func (s *SkipList) Set(k, v interface{}) {
	h, p := hash(k), unsafe.Pointer(&v)
	e := &entry{k: k, p: p, hash: h, next: make([]*entry, s.random())}
	if len(s.headers) == 0 {
		s.headers = make([]*entry, e.level()+1)
		for i := 0; i < s.random(); i++ {
			s.headers[i] = e
		}
		return
	}
	level := s.level()

	next := s.headers[level]
	for next == nil && level > 0 {
		level--
		next = s.headers[level]
	}
	var prev *entry
	var pres = make([]*entry, e.level()+1)
	for next != nil {
		if next.k == k {
			next.p = e.p
			return
		}
		if h > next.hash {
			if level < len(pres) {
				pres[level] = prev
			}
			if level > 0 {
				level--
				if prev == nil {
					next = s.headers[level]
				} else {
					next = prev.next[level]
				}
				continue
			}
			for i := 0; i < len(pres); i++ {
				if pres[i] == nil {
					e.next[i] = s.headers[i]
					s.headers[i] = e
				} else {
					e.next[i] = pres[i].next[i]
					pres[i].next[i] = e
				}
			}
			return
		}
		prev = next
		next = next.next[level]
	}
}

func (s SkipList) Get(k interface{}) (interface{}, bool) {
	if len(s.headers) == 0 {
		return nil, false
	}
	h := hash(k)
	level := s.level()
	next := s.headers[level]
	for next == nil && level > 0 {
		level--
		next = s.headers[level]
	}
	var prev *entry
	for next != nil {
		if next.k == k {
			return next.value(), true
		}
		if h > next.hash {
			if level == 0 {
				return nil, false
			}
			level--
			if prev == nil {
				next = s.headers[level]
			} else {
				next = prev.next[level]
			}
		}
		prev = next
		next = next.next[level]
	}
	return nil, false
}

func (s *SkipList) Del(k interface{}) bool {
	if len(s.headers) == 0 {
		return false
	}
	h := hash(k)
	level := s.level()
	next := s.headers[level]
	for next == nil && level > 0 {
		level--
		next = s.headers[level]
	}
	var prev *entry
	var pres = make([]*entry, s.level()+1)
	for next != nil {
		if next.k == k && level == 0 {
			pres[level] = prev
			for i := 0; i <= next.level(); i++ {
				if pres[i] == nil {
					s.headers[i] = next.next[i]
				} else {
					pres[i] = next.next[i]
				}
			}
			return true
		}
		if next.k == k || h > next.hash {
			if h > next.hash && level == 0 {
				return false
			}
			pres[level] = prev
			level--
			if prev == nil {
				next = s.headers[level]
			} else {
				next = prev.next[level]
			}
			continue
		}
		prev = next
		next = next.next[level]
	}
	return false
}

func (s SkipList) print() string {
	buf := bytes.Buffer{}
	for i := s.level(); i >= 0; i-- {
		buf.WriteString("||->")
		next := s.headers[i]
		for next != nil {
			buf.WriteString(fmt.Sprintf("%v", next.value()))
			buf.WriteString("->")
			next = next.next[i]
		}
		buf.WriteString("\n")
	}
	return buf.String()
}
