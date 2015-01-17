package test

import (
	"github.com/albrow/go-pascal/common"
	"github.com/albrow/go-pascal/implementations"
	"testing"
)

func BenchmarkNaive(b *testing.B) {
	benchmarkPascaler(b, implementations.Naive)
}

func BenchmarkRecursive(b *testing.B) {
	// The recursive implementation is much, much slower for large n and m.
	// With the parameters we are using to benchmark the other two implementations,
	// the recursive one will simply timeout. So we'll skip it for now. If you'd like
	// to see how slow it is, reduce n and m to something like (30, 15).
	b.Skip("Skipping much slower recursive implementation")
	benchmarkPascaler(b, implementations.Recursive)
}

func BenchmarkBuiltin(b *testing.B) {
	benchmarkPascaler(b, implementations.Builtin)
}

func benchmarkPascaler(b *testing.B, p common.Pascaler) {
	for i := 0; i < b.N; i++ {
		p.Pascal(2e4, 1e3)
	}
}
