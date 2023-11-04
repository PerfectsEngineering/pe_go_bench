package random

import "testing"

func BenchmarkGenerateRandomNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandomNumber()
	}
}
