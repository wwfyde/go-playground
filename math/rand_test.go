package math

import (
	"math/rand"
	"testing"
)

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {

		rand.Int()
		//fmt.Println(result)
	}
}
