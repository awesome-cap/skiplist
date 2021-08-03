# SkipList
Fast, thread-safe skiplist.
# Usage
Install
```go
go get -u github.com/awesome-cap/skiplist
```
Using
```go
import "github.com/awesome-cap/skiplist"

func main() {
    s := skiplist.New(18)
    s.Set(1, 1)
    s.Get(1)
    s.Del(1)
}
```
# Benchmark
```text
goos: windows
goarch: amd64
pkg: github.com/awesome-cap/skiplist
cpu: 11th Gen Intel(R) Core(TM) i5-11400 @ 2.60GHz
BenchmarkSet_SkipList-12                                 5123154               263.4 ns/op
BenchmarkSet_JumpList-12                                 5811688               212.6 ns/op
BenchmarkSet_FastList-12                                 5783910               224.7 ns/op
BenchmarkGet_SkipList-12                                 9628252               144.0 ns/op
BenchmarkGet_JumpList-12                                 9546835               135.9 ns/op
BenchmarkGet_FastList-12                                 9474108               145.0 ns/op
BenchmarkSetRandom_SkipList-12                           1000000              1639 ns/op
BenchmarkSetRandom_JumpList-12                           1000000              1341 ns/op
BenchmarkSetRandom_FastList-12                           1000000              1319 ns/op
BenchmarkGetRandom_SkipList-12                           1000000              1840 ns/op
BenchmarkGetRandom_JumpList-12                           1000000              1629 ns/op
BenchmarkGetRandom_FastList-12                           1000000              1675 ns/op
BenchmarkSetAndGet_SkipList-12                           3379516               350.9 ns/op
BenchmarkSetAndGet_JumpList-12                           3781561               350.0 ns/op
BenchmarkSetAndGet_FastList-12                           3634700               356.4 ns/op
BenchmarkSetRandomAndGetRandom_SkipList-12               1000000              3648 ns/op
BenchmarkSetRandomAndGetRandom_JumpList-12               1000000              3111 ns/op
BenchmarkSetRandomAndGetRandom_FastList-12               1000000              3088 ns/op
BenchmarkSetAndGetAsync_SkipList-12                      4360150               261.2 ns/op
BenchmarkSetAndGetAsync_JumpList-12                      4163358               323.2 ns/op
BenchmarkSetAndGetAsync_FastList-12                      3944946               333.5 ns/op
BenchmarkSetRandomAndGetRandomAsync_SkipList-12          1000000              1598 ns/op
BenchmarkSetRandomAndGetRandomAsync_JumpList-12          1000000              2911 ns/op
BenchmarkSetRandomAndGetRandomAsync_FastList-12          1000000              3230 ns/op
BenchmarkSetParallel_SkipList-12                           10000           1835569 ns/op
BenchmarkSetParallel_JumpList-12                           10000           1590117 ns/op
BenchmarkSetParallel_FastList-12                           10000           1543311 ns/op
```