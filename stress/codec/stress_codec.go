package main

import (
	"flag"
	"fmt"
	"github.com/10cella/yomo-y3-stress-testing/internal/tester"
	"io"
	"os"
	"time"
)

func main() {
	file := flag.String("file", "./assets/codec/codec_test_data.bin", "file path of codec data")
	key := flag.Int("key", 0x02, "key of observe")
	target := flag.Int64("target", 1000, "target for reporter")
	flag.Parse()

	run(1, byte(*key), *target, *file)

	time.Sleep(3600 * time.Second)
}

func run(index int64, key byte, target int64, file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	r := io.Reader(f)

	startTime := time.Now().UnixNano()
	w := tester.NewCodecTester(key, target, func(endTime int64) {
		elapsedTime := endTime - startTime // ns
		fmt.Printf("codec decode[%v]: key=%#x counter=%v, elapsedTime=%vms, singleTime=%vns \n",
			index, key, target, elapsedTime/1e6, elapsedTime/target)

	})
	io.Copy(w, r)

	_ = f.Close()
}
