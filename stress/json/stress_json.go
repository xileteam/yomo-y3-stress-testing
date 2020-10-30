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
	file := flag.String("file", "./assets/json/json_test_data.bin", "file path of json data")
	key := flag.String("key", "k1", "key of observe")
	target := flag.Int64("target", 1000, "target for reporter")
	flag.Parse()

	go run(1, *key, *target, *file)

	time.Sleep(3600 * time.Second)
}

func run(index int64, key string, target int64, file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	r := io.Reader(f)

	startTime := time.Now().UnixNano()
	w := tester.NewJsonTester(key, target, func(endTime int64) {
		elapsedTime := endTime - startTime // ns
		fmt.Printf("json decode[%v]: key=%v counter=%v, elapsedTime=%vms, singleTime=%vns \n",
			index, key, target, elapsedTime/1e6, elapsedTime/target)
	})
	io.Copy(w, r)

	_ = f.Close()
}
