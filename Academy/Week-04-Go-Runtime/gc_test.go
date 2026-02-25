package main

import (
	"testing"
)

type BigStruct struct {
	Data [1024]int
}

// BenchmarkBad: Allocates a new object every time
func BenchmarkBad(b *testing.B) {
	var s *BigStruct
	for i := 0; i < b.N; i++ {
		s = &BigStruct{} // Forced heap allocation
		_ = s
	}
}

// BenchmarkGood: Reuses the same object (simulating a pool or stack)
func BenchmarkGood(b *testing.B) {
	s := &BigStruct{}
	for i := 0; i < b.N; i++ {
		// Just reset the data instead of allocating new memory
		for j := 0; j < 1024; j++ {
			s.Data[j] = 0
		}
		_ = s
	}
}
