package decoder

import (
	"errors"
	"github.com/10cella/yomo-y3-stress-testing/internal/generator"
	"testing"
)

func Benchmark_Json_k1(b *testing.B) {
	key := "k1"
	data := generator.NewJsonTestData().GenData()
	data = append(data, TokenEnd)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if TakeValueFromJson(key, data) == nil {
			panic(errors.New("take is failure"))
		}
	}
}

func Benchmark_Json_k32(b *testing.B) {
	key := "k32"
	data := generator.NewJsonTestData().GenData()
	data = append(data, TokenEnd)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if TakeValueFromJson(key, data) == nil {
			panic(errors.New("take is failure"))
		}
	}
}

func Benchmark_Json_k63(b *testing.B) {
	key := "k63"
	data := generator.NewJsonTestData().GenData()
	data = append(data, TokenEnd)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if TakeValueFromJson(key, data) == nil {
			panic(errors.New("take is failure"))
		}
	}
}

