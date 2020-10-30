package decoder

import (
	"errors"
	"github.com/10cella/yomo-y3-stress-testing/internal/generator"
	"testing"
)

func Benchmark_Codec_0x01(b *testing.B) {
	var key byte = 0x01
	data := generator.NewCodecTestData().GenData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if TakeValueFromCodec(key, data) == nil {
			panic(errors.New("take is failure"))
		}
	}
}

func Benchmark_Codec_0x20(b *testing.B) {
	var key byte = 0x20
	data := generator.NewCodecTestData().GenData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if TakeValueFromCodec(key, data) == nil {
			panic(errors.New("take is failure"))
		}
	}
}

func Benchmark_Codec_0x3f(b *testing.B) {
	var key byte = 0x3f
	data := generator.NewCodecTestData().GenData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if TakeValueFromCodec(key, data) == nil {
			panic(errors.New("take is failure"))
		}
	}
}
