package main

import (
	"flag"
	"fmt"
	"github.com/10cella/yomo-y3-stress-testing/internal/generator"
	"os"
)

func main() {
	file := flag.String("file", "./assets/json/json_test_data.bin", "file path of test data")
	count := flag.Int("count", 10, "count of data")

	f, _ := os.Create(*file)
	for i := 1; i <= *count; i++ {
		_, _ = f.Write(generator.NewJsonTestData().GenData())
		_, _ = f.Write([]byte("\n"))
	}
	_ = f.Close()

	fmt.Printf("write data to file successfully, file=%v, count=%v\n", *file, *count)
}
