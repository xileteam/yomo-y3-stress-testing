package report_parallel

import (
	"errors"
	"github.com/10cella/yomo-y3-stress-testing/internal/decoder"
	"github.com/10cella/yomo-y3-stress-testing/internal/generator"
	"testing"
)

func Benchmark_Codec_C03_K02(b *testing.B) {
	var key byte = 0x02
	data := generator.NewCodecTestData().GenDataBy(3)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			if decoder.TakeValueFromCodec(key, data) == nil {
				panic(errors.New("take is failure"))
			}
		}
	})
}

func Benchmark_Json_C03_K02(b *testing.B) {
	key := "k2"
	data := generator.NewJsonTestData().GenDataBy(3)
	data = append(data, decoder.TokenEnd)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			if decoder.TakeValueFromJson(key, data) == nil {
				panic(errors.New("take is failure"))
			}
		}
	})
}

func Benchmark_Codec_C16_K08(b *testing.B) {
	var key byte = 0x08
	data := generator.NewCodecTestData().GenDataBy(16)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			if decoder.TakeValueFromCodec(key, data) == nil {
				panic(errors.New("take is failure"))
			}
		}
	})
}

func Benchmark_Json_C16_K08(b *testing.B) {
	key := "k8"
	data := generator.NewJsonTestData().GenDataBy(16)
	data = append(data, decoder.TokenEnd)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			if decoder.TakeValueFromJson(key, data) == nil {
				panic(errors.New("take is failure"))
			}
		}
	})
}

func Benchmark_Codec_C32_K16(b *testing.B) {
	var key byte = 0x10
	data := generator.NewCodecTestData().GenDataBy(32)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			if decoder.TakeValueFromCodec(key, data) == nil {
				panic(errors.New("take is failure"))
			}
		}
	})
}

func Benchmark_Json_C32_K16(b *testing.B) {
	key := "k16"
	data := generator.NewJsonTestData().GenDataBy(32)
	data = append(data, decoder.TokenEnd)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			if decoder.TakeValueFromJson(key, data) == nil {
				panic(errors.New("take is failure"))
			}
		}
	})
}

func Benchmark_Codec_C63_K32(b *testing.B) {
	var key byte = 0x20
	data := generator.NewCodecTestData().GenDataBy(63)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			if decoder.TakeValueFromCodec(key, data) == nil {
				panic(errors.New("take is failure"))
			}
		}
	})
}

func Benchmark_Json_C63_K32(b *testing.B) {
	key := "k32"
	data := generator.NewJsonTestData().GenDataBy(63)
	data = append(data, decoder.TokenEnd)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			if decoder.TakeValueFromJson(key, data) == nil {
				panic(errors.New("take is failure"))
			}
		}
	})
}
