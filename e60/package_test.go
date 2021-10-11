package main

import "testing"

func TestConc1(t *testing.T) {
	if conc(1, 2) != 12 {
		t.Error("bad result", 1, 2)
	}
}

func TestConc2(t *testing.T) {
	if conc(12, 45) != 1245 {
		t.Error("bad result", 12, 45)
	}
}
func TestConc3(t *testing.T) {
	if conc(1, 1000) != 11000 {
		t.Error("bad result", 1, 1000, conc(1, 1000))
	}
}

func TestConc4(t *testing.T) {
	if conc(990, 42) != 99042 {
		t.Error("bad result", 990, 42, conc(990, 42))
	}
}

func BenchmarkConc1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conc1(12, 45)
	}
}

func BenchmarkConc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conc(12, 45)
	}
}
