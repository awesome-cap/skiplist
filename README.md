# SkipList
Atomic skiplist.
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
BenchmarkSkipList_Set-12                                 4496858               297.4 ns/op
BenchmarkJumpList_Set-12                                 5812356               215.1 ns/op
BenchmarkSkipList_SetRandom-12                           1000000              1790 ns/op
BenchmarkJumpList_SetRandom-12                           1000000              1416 ns/op
BenchmarkSkipList_SetAndGet-12                           3065451               402.7 ns/op
BenchmarkJumpList_SetAndGet-12                           3624136               358.6 ns/op
BenchmarkSkipList_SetRandomAndGetRandom-12               1000000              4129 ns/op
BenchmarkJumpList_SetRandomAndGetRandom-12               1000000              3348 ns/op
BenchmarkSkipList_SetAndGetAsync-12                      3992883               300.6 ns/op
BenchmarkJumpList_SetAndGetAsync-12                      3891741               323.9 ns/op
BenchmarkSkipList_SetRandomAndGetRandomAsync-12          1000000              2143 ns/op
BenchmarkJumpList_SetRandomAndGetRandomAsync-12          1000000              3231 ns/op
BenchmarkSkipList_SetParallel-12                           10000           1922838 ns/op
BenchmarkJumpList_SetParallel-12                           10000           1697410 ns/op
```