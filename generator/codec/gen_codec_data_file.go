package main

import (
	"flag"
	"fmt"
	"github.com/10cella/yomo-y3-stress-testing/internal/generator"
	"io/ioutil"
)

func main() {
	file := flag.String("file", "./assets/codec/codec_test_data.bin", "file path of test data")

	data := generator.NewCodecTestData().GenData()
	fmt.Printf("%#x\n", data)

	err := ioutil.WriteFile(*file, data, 0644) // -rw-r--r--
	if err != nil {
		panic(err)
	}

	fmt.Printf("write data to file successfully, file=%v\n", *file)
}
