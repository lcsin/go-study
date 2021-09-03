package fib

import "testing"

// 性能比较函数
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B) {
	benchmarkFib(b, 10)
}

func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b, 100)
}

func BenchmarkFib4(b *testing.B) {
	benchmarkFib(b, 1000)
}

func BenchmarkFib8(b *testing.B) {
	benchmarkFib(b, 10000)
}

func BenchmarkFib10(b *testing.B) {
	benchmarkFib(b, 100000)
}
