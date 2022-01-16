package popcount

import (
	"testing"
)

var testX uint64 = 100000

func BenchmarkPopCount(b *testing.B) {
	var i int
	for i = 0; i <= b.N; i++ {
		_ = PopCount(testX)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	var i int
	for i = 0; i <= b.N; i++ {
		_ = PopCount2(testX)
	}
}

func BenchmarkPopCount3(b *testing.B) {
	var i int
	for i = 0; i <= b.N; i++ {
		_ = PopCount3(testX)
	}
}

func BenchmarkPopCount4(b *testing.B) {
	var i int
	for i = 0; i <= b.N; i++ {
		_ = PopCount4(testX)
	}
}
