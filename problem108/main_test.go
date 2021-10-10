package main

import "testing"

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(1000)
	}
}

func BenchmarkSolve2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(1000)
	}
}
