package skiplist

import (
	"github.com/abbychau/jumplist"
	"math/rand"
	"sync"
	"testing"
)

var (
	arr   []int
	batch = 1000 * 10000
)

func init() {
	for i := 0; i < batch; i++ {
		arr = append(arr, rand.Int())
	}
}

type tester interface {
	Set(b *testing.B)
	Get(b *testing.B)
	SetRand(b *testing.B)
	GetRand(b *testing.B)
}

func newSkipListTester() tester {
	return &skipListTester{s: New(18)}
}

func newJumpListTester() tester {
	return &jumpListTester{s: jumplist.New()}
}

type skipListTester struct {
	s *SkipList
}

func (s *skipListTester) Set(b *testing.B) {
	for j := 0; j < b.N; j++ {
		s.s.Set(float64(j), j)
	}
}

func (s *skipListTester) Get(b *testing.B) {
	for j := 0; j < b.N; j++ {
		s.s.Get(float64(j))
	}
}

func (s *skipListTester) SetRand(b *testing.B) {
	for j := 0; j < b.N; j++ {
		s.s.Set(float64(arr[j]), arr[j])
	}
}

func (s *skipListTester) GetRand(b *testing.B) {
	for j := 0; j < b.N; j++ {
		s.s.Get(float64(arr[j]))
	}
}

type jumpListTester struct {
	s *jumplist.SkipList
}

func (s *jumpListTester) Set(b *testing.B) {
	for j := 0; j < b.N; j++ {
		s.s.Set(float64(j), j)
	}
}

func (s *jumpListTester) Get(b *testing.B) {
	for j := 0; j < b.N; j++ {
		s.s.Get(float64(j))
	}
}

func (s *jumpListTester) SetRand(b *testing.B) {
	for j := 0; j < b.N; j++ {
		s.s.Set(float64(arr[j]), arr[j])
	}
}

func (s *jumpListTester) GetRand(b *testing.B) {
	for j := 0; j < b.N; j++ {
		s.s.Get(float64(arr[j]))
	}
}

func testSet(b *testing.B, t tester) {
	b.ResetTimer()
	t.Set(b)
}

func testSetRandom(b *testing.B, t tester) {
	b.ResetTimer()
	t.SetRand(b)
}

func testSetAndGet(b *testing.B, t tester) {
	b.ResetTimer()
	t.Set(b)
	t.Get(b)
}

func testSetRandomAndGetRandom(b *testing.B, t tester) {
	b.ResetTimer()
	t.SetRand(b)
	t.GetRand(b)
}

func testSetAndGetAsync(b *testing.B, t tester) {
	b.ResetTimer()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		t.Set(b)
		wg.Done()
	}()
	go func() {
		t.Get(b)
		wg.Done()
	}()
	wg.Wait()
}

func testSetRandomAndGetRandomAsync(b *testing.B, t tester) {
	b.ResetTimer()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		t.SetRand(b)
		wg.Done()
	}()
	go func() {
		t.GetRand(b)
		wg.Done()
	}()
	wg.Wait()
}

func testSetParallel(b *testing.B, t tester) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t.Set(b)
		}
	})
}

func BenchmarkSkipList_Set(b *testing.B) { testSet(b, newSkipListTester()) }
func BenchmarkJumpList_Set(b *testing.B) { testSet(b, newJumpListTester()) }

func BenchmarkSkipList_SetRandom(b *testing.B) { testSetRandom(b, newSkipListTester()) }
func BenchmarkJumpList_SetRandom(b *testing.B) { testSetRandom(b, newJumpListTester()) }

func BenchmarkSkipList_SetAndGet(b *testing.B) { testSetAndGet(b, newSkipListTester()) }
func BenchmarkJumpList_SetAndGet(b *testing.B) { testSetAndGet(b, newJumpListTester()) }

func BenchmarkSkipList_SetRandomAndGetRandom(b *testing.B) {
	testSetRandomAndGetRandom(b, newSkipListTester())
}
func BenchmarkJumpList_SetRandomAndGetRandom(b *testing.B) {
	testSetRandomAndGetRandom(b, newJumpListTester())
}

func BenchmarkSkipList_SetAndGetAsync(b *testing.B) { testSetAndGetAsync(b, newSkipListTester()) }
func BenchmarkJumpList_SetAndGetAsync(b *testing.B) { testSetAndGetAsync(b, newJumpListTester()) }

func BenchmarkSkipList_SetRandomAndGetRandomAsync(b *testing.B) {
	testSetRandomAndGetRandomAsync(b, newSkipListTester())
}
func BenchmarkJumpList_SetRandomAndGetRandomAsync(b *testing.B) {
	testSetRandomAndGetRandomAsync(b, newJumpListTester())
}

func BenchmarkSkipList_SetParallel(b *testing.B) { testSetParallel(b, newSkipListTester()) }
func BenchmarkJumpList_SetParallel(b *testing.B) { testSetParallel(b, newJumpListTester()) }
