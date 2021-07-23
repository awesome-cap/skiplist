package skiplist

import (
	"fmt"
	"testing"
)

func TestSkipList_Set(t *testing.T) {
	s := New(48)
	batch := 10
	for i := 0; i < batch; i++ {
		s.Set(i, i)
	}
	fmt.Println(s.print())
	for i := 0; i < batch; i++ {
		fmt.Println(s.Get(i))
	}

	s.Del(6)
	fmt.Println(s.print())
}
