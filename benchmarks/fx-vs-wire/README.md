# Google Wire vs Uber Fx Benchmark

This benchmark reproduces and validates the performance difference between **Google Wire** (compile-time code generation) and **Uber Fx** (runtime reflection DI framework).

---

## 📊 Benchmark Results (Apple Silicon M-Series)

```bash
$ go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: github.com/skyrocket-qy/fx-vs-wire
BenchmarkWire-10         	1000000000	         0.2271 ns/op	       0 B/op	       0 allocs/op
BenchmarkDirect-10       	1000000000	         0.2254 ns/op	       0 B/op	       0 allocs/op
BenchmarkFx-10           	   39981	     30595 ns/op	   35621 B/op	     511 allocs/op
BenchmarkFxNewOnly-10    	   43504	     27659 ns/op	   35355 B/op	     499 allocs/op
```

### Key Insights:
1. **Google Wire vs Direct:** Wire achieves the exact same performance as hand-written direct instantiation (`~0.22 ns/op`, `0 B/op`, `0 allocs/op`) because it generates static, inlinable Go code.
2. **Uber Fx Reflection Cost:** Constructing the Fx container requires runtime reflection and dependency graph building, incurring `~30 µs`, `35.6 KB` of heap allocations, and `~500 allocs/op`.

---

## 🚀 How to Run

### 1. Run the Benchmarks
```bash
go test -bench=. -benchmem
```

### 2. Regenerate Wire Code
If you modify `wire.go`:
```bash
go run github.com/google/wire/cmd/wire@latest .
```
