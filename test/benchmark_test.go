package test

import "testing"

// go test -v -bench=Hello
func Benchmark_Hello(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
	}

	b.StopTimer()
}
