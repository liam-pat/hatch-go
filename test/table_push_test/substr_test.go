package main

import "testing"

func lengthOfNonRepeatingSubString(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if existIndex, ok := lastOccurred[ch]; ok && existIndex >= start {
			start = existIndex + 1
		}
		// 当前位置减去开始的位置再加上 1.若比 maxLength，那么字符串就有另外一个比第一个长，
		if i-start+1 >= maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

// 简单的表格驱动测试
func TestSubstr(t *testing.T) {
	tests := []struct {
		s      string
		result int
	}{
		{"abcdabe", 5},
		{"pwwkew", 3},
		{"", 0},
		{"b", 1},
		{"我在测试测试测试", 4},
		{"黑灰化肥灰会挥发发灰黑讳为黑灰花会飞", 7},
	}

	for _, tt := range tests {
		actualResult := lengthOfNonRepeatingSubString(tt.s)
		if actualResult != tt.result {
			t.Errorf("Got %d for input: [ %s ]; expected: %d", actualResult, tt.s, tt.result)
		}
	}
}

//cpu 性能测试
func BenchmarkSubstr(b *testing.B) {
	s := "黑灰化肥灰会挥发发灰黑讳为黑灰花会飞"
	result := 7

	for i := 0; i <= b.N; i++ {
		actualResult := lengthOfNonRepeatingSubString(s)

		if result != actualResult {
			b.Errorf("Got %d for input: [ %s ]; expected: %d", actualResult, s, result)
		}
	}
}
