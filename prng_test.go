package randutil

import (
	"math/rand"
	"testing"
)

func BenchmarkUniformUint32(b *testing.B) {
	for i := uint32(0); i < uint32(b.N); i++ {
		UniformUint32(i)
	}
}

func BenchmarkUniformInt32Range(b *testing.B) {
	for i := int64(0); i < int64(b.N); i++ {
		UniformInt32Range(i, 0, 1e5)
	}
}

func BenchmarkUniformFloat64(b *testing.B) {
	for i := int64(0); i < int64(b.N); i++ {
		UniformFloat64(i)
	}
}

func BenchmarkNormalUint32(b *testing.B) {
	for i := uint32(0); i < uint32(b.N); i++ {
		NormalUint32(i)
	}
}

func BenchmarkNormalInt32Range(b *testing.B) {
	for i := int64(0); i < int64(b.N); i++ {
		NormalInt32Range(i, 0, 1e5)
	}
}

func BenchmarkNormalFloat64(b *testing.B) {
	for i := int64(0); i < int64(b.N); i++ {
		NormalFloat64(i)
	}
}

func BenchmarkStandardInt63(b *testing.B) {
	rnum := rand.New(rand.NewSource(0))
	for i := 0; i < b.N; i++ {
		rnum.Int63()
	}
}

func BenchmarkStandardNormFloat64(b *testing.B) {
	rnum := rand.New(rand.NewSource(0))
	for i := 0; i < b.N; i++ {
		rnum.NormFloat64()
	}
}

func BenchmarkStandardInt63n(b *testing.B) {
	rnum := rand.New(rand.NewSource(0))
	for i := 0; i < b.N; i++ {
		rnum.Int63n(1e5)
	}
}
