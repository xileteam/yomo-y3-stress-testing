package tester

import (
	"github.com/10cella/yomo-y3-stress-testing/internal/decoder"
	y3 "github.com/yomorun/yomo-codec-golang"
	"io"
	"sync/atomic"
	"time"
)

type codecTester struct {
	io.Writer
	key byte

	Tag       *decoder.Tag
	Sbuf      []byte
	Size      int32
	Length    int32
	LengthBuf []byte

	//Values  []interface{}
	counter int64
	target  int64

	report func(endTime int64)
}

func NewCodecTester(observe byte, target int64, report func(endTime int64)) *codecTester {
	return &codecTester{
		key: observe,

		Tag:       nil,
		Sbuf:      make([]byte, 0),
		Size:      0,
		Length:    0,
		LengthBuf: make([]byte, 0),

		//Values:  make([]interface{}, 0),
		counter: 0,
		target:  target,
		report:  report,
	}
}

func (w *codecTester) Write(buf []byte) (int, error) {
	w.decode(buf)
	return len(buf), nil
}

func (w *codecTester) decode(buf []byte) {
	//fmt.Printf("#301 buf=%s\n", codes.FormatBytes(buf))
	for _, c := range buf {
		// tag
		if w.Tag == nil {
			w.Tag = decoder.NewTag(c)
			w.Sbuf = append(w.Sbuf, c)
			continue
		}

		// length
		if w.Size == 0 {
			w.LengthBuf = append(w.LengthBuf, c)
			w.Sbuf = append(w.Sbuf, c)
			length, size, err := decoder.DecodeLength(w.LengthBuf)
			if err != nil {
				continue
			}
			w.Length = length
			w.Size = size
			continue
		}

		w.Sbuf = append(w.Sbuf, c)

		// not the current tag
		if w.key != w.Tag.SeqID() {
			//fmt.Printf("#301 Tag.SeqID()=%#x, len(buf)=%v, int(1+w.Size+w.Length)=%v\n",
			//	w.Tag.SeqID(), len(buf), int(1+w.Size+w.Length))
			if len(w.Sbuf) == int(1+w.Size+w.Length){
				//value := decoder.TakeValueFromCodec(w.key, w.Sbuf)
				//w.Values = append(w.Values, value)
				if decoder.TakeValueFromCodec(w.key, w.Sbuf) != nil {
					w.trigger()
				}
				w.reset()
				continue
			}
		}

		// current tag is ending
		if int32(len(w.Sbuf)) == 1+w.Size+w.Length {
			if w.Tag.IsNode() {
				packet, _, err := y3.DecodeNodePacket(w.Sbuf)
				if err != nil {
					w.endLoop(false)
					continue
				}
				if ok, _, _ := decoder.MatchingKey(w.key, packet); ok {
					//w.Values = append(w.Values, p)
					//w.endLoop(true)
					w.trigger()
					w.reset()
					continue
				}
			} else {
				_, _, _, err := y3.DecodePrimitivePacket(w.Sbuf)
				if err != nil {
					w.endLoop(false)
					continue
				}
				//w.Values = append(w.Values, *packet)
				//w.endLoop(true)
				w.trigger()
				w.reset()
				continue
			}
		}

	}
}

func (w *codecTester) endLoop(isSuccessful bool) {
	if isSuccessful {
		w.trigger()
	}
	w.reset()
}

func (w *codecTester)  trigger() {
	count := atomic.AddInt64(&w.counter, 1)
	if count == w.target {
		w.report(time.Now().UnixNano())
	}
}

func (w *codecTester) reset() {
	w.Tag = nil
	w.LengthBuf = make([]byte, 0)
	w.Length = 0
	w.Size = 0
	w.Sbuf = make([]byte, 0)
}
