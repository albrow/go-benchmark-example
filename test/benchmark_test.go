package test

import (
	"github.com/albrow/go-pascal/common"
	"github.com/albrow/go-pascal/implementations"
	"testing"
)

func BenchmarkNaive(b *testing.B) {
	benchmarkPascaler(b, implementations.Naive)
}

func benchmarkPascaler(b *testing.B, p common.Pascaler) {
	for i := 0; i < b.N; i++ {
		p.Pascal(100, 100)
	}
}
