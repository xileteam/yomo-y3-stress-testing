package tester

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/10cella/yomo-y3-stress-testing/internal/decoder"
	"io"
	"sync/atomic"
	"time"
)

type jsonTester struct {
	io.Writer

	key  string
	sbuf []byte

	counter int64
	target int64

	report func(endTime int64)
}

func NewJsonTester(observe string, target int64, report func(endTime int64)) *jsonTester {
	return &jsonTester{
		key:  observe,
		sbuf: make([]byte, 0),
		counter: 0,
		target: target,
		report: report,
	}
}

func (w *jsonTester) Write(buf []byte) (int, error) {
	w.decode(buf)
	return len(buf), nil
}

func (w *jsonTester) decode(buf []byte) {
	for _, c := range buf {
		if c == decoder.TokenEnd {
			var m map[string]interface{}
			if err := json.Unmarshal(w.sbuf, &m); err != nil {
				panic(errors.New(fmt.Sprintf("data is not json: %#x\n", buf)))
			}

			for k, v := range m {
				if k == w.key {
					if v != nil{
						count := atomic.AddInt64(&w.counter, 1)
						if count == w.target {
							w.report(time.Now().UnixNano())
						}
					}
					break
				}
			}

			w.sbuf = make([]byte, 0)
		}

		w.sbuf = append(w.sbuf, c)
	}
}
