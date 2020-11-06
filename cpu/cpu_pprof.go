package main

import (
	"errors"
	"fmt"
	"github.com/10cella/yomo-y3-stress-testing/internal/decoder"
	"github.com/10cella/yomo-y3-stress-testing/internal/generator"
	"github.com/10cella/yomo-y3-stress-testing/internal/pprof"
	"time"
)

func main() {
	dataCodec := generator.NewCodecTestData().GenDataBy(63)
	dataJson := generator.NewJsonTestData().GenDataBy(63)
	dataJson = append(dataJson, decoder.TokenEnd)

	// pprof
	fmt.Printf("start pprof\n")
	go pprof.Run()
	time.Sleep(5 * time.Second)
	
	fmt.Printf("start testing...\n")
	for {
		if decoder.TakeValueFromCodec(0x20, dataCodec) == nil {
			panic(errors.New("take is failure"))
		}
		if decoder.TakeValueFromJson("k32", dataJson) == nil {
			panic(errors.New("take is failure"))
		}
		//time.Sleep(200 * time.Microsecond)
	}

	time.Sleep(3600 * time.Second)
}



