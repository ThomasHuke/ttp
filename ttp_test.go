package ttp

import (
	"testing"
)

func TestTTP(t *testing.T) {
	TT()
}

func BenchmarkTT(t *testing.B) {
	for i := 0; i < t.N; i++ {
		TT()

	}
}

// go test -bench=.
// goos: darwin
// goarch: amd64
// pkg: github.com/ThomasHuke/ttp
// BenchmarkTT-4   	10000000	       141 ns/op
// PASS
// ok  	github.com/ThomasHuke/ttp	1.572s
