package main

import "testing"

func BenchmarkValidFloat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		validFloat(10, 22, 31)
	}
}

func BenchmarkValidInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		validInt(10, 22, 31)
	}
}
