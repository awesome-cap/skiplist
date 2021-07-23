package skiplist

import (
	"bytes"
	"fmt"
	"math/rand"
	"unsafe"
)

type SkipList struct {
	limit   int
	level   int
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

func (s SkipList) random() int {
	i := 0
	for i < s.limit && rand.Float32() <= 0.25 {
		i++
	}
	return i
}

func (s *SkipList) Set(k, v interface{}) {
	h, p := hash(k), unsafe.Pointer(&v)
	e := &entry{k: k, p: p, hash: h, next: make([]*entry, s.random()+1)}
	if len(s.headers) == 0 || (s.level == 0 && s.headers[s.level] == nil) {
		if len(s.headers) == 0 {
			s.headers = make([]*entry, e.level()+1)
		}
		for i := 0; i <= e.level(); i++ {
			s.headers[i] = e
		}
		s.level = max(s.level, e.level())
		return
	}
	level := s.level
	next := s.headers[level]
	var prev *entry
	var pres = make([]*entry, e.level()+1)
	for {
		if next == nil {
			if level < len(pres) {
				pres[level] = prev
			}
			if level > 0 {
				level--
				if prev != nil {
					next = prev.next[level]
				}
				continue
			}
			break
		}
		if next.k == k {
			next.p = e.p
			return
		}
		if h < next.hash {
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
			break
		}
		prev = next
		next = next.next[level]
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
	s.level = max(s.level, e.level())
}

func (s SkipList) Get(k interface{}) (interface{}, bool) {
	if len(s.headers) == 0 {
		return nil, false
	}
	h := hash(k)
	level := s.level
	next := s.headers[level]
	var prev *entry
	for {
		if next == nil {
			if level > 0 {
				level--
				if prev != nil {
					next = prev.next[level]
				}
				continue
			}
			break
		}
		if next.k == k {
			return next.value(), true
		}
		if h < next.hash {
			if level == 0 {
				return nil, false
			}
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
	return nil, false
}

func (s *SkipList) Del(k interface{}) bool {
	if len(s.headers) == 0 {
		return false
	}
	h := hash(k)
	level := s.level
	next := s.headers[level]
	var prev *entry
	var pres = make([]*entry, s.level+1)
	for {
		if next == nil {
			if level < len(pres) {
				pres[level] = prev
			}
			if level > 0 {
				level--
				if prev != nil {
					next = prev.next[level]
				}
				continue
			}
			break
		}
		if next.k == k && level == 0 {
			pres[level] = prev
			for i := 0; i <= next.level(); i++ {
				if pres[i] == nil {
					s.headers[i] = next.next[i]
				} else {
					pres[i] = next.next[i]
				}
			}
			// skip list level
			l := s.level
			for ; l >= 0 && s.headers[l] == nil; l-- {
			}
			s.level = l
			return true
		}
		if next.k == k || h < next.hash {
			if h < next.hash && level == 0 {
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
	for i := s.level; i >= 0; i-- {
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

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
