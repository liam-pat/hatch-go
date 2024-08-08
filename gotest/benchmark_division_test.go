package gotest

import (
	"errors"
	"testing"
)

// go test benchmark_division_test.go -test.bench=".*"

func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Division(4, 5)

	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Division(4, 5)
	}
}

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}

	return a / b, nil
}
