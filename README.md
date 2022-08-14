# concurrent-map-sample

Performance comparison between [orcaman/concurrent-map](https://github.com/orcaman/concurrent-map) and sync.Map

## benchmark test result

```
$ go test -bench .
goos: linux
goarch: amd64
pkg: concurrent-map-sample
cpu: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz
BenchmarkSyncMap-4                667300              2510 ns/op             229 B/op          7 allocs/op
BenchmarkConcurrentMap-4         1000000              1392 ns/op             226 B/op          2 allocs/op
PASS
ok      concurrent-map-sample   3.124s
```

[orcaman/concurrent-map/v2](https://github.com/orcaman/concurrent-map) performs better than sync.Map

Because orcaman/concurrent-map uses the sync/atomic package internally and sync.Map uses sync.Mutex internally to ensure thread-safety.
Unlike sync.Mutex, sync.RWMutex, and sync.Cond, the sync/atomic package does not use context switches internally, so the load is low.
