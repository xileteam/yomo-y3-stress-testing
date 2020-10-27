package generator

import (
	"errors"
	"github.com/10cella/yomo-y3-stress-testing/internal/mapping"
)

type codecTestData struct {
	DefaultCount uint
}

func NewCodecTestData() *codecTestData {
	return &codecTestData{DefaultCount: 63}
}

func (c *codecTestData) GenData() []byte {
	return c.GenDataBy(c.DefaultCount)
}

func (c *codecTestData) GenDataBy(count uint) []byte {
	if count > c.DefaultCount {
		panic(errors.New("out of bounds"))
	}

	section := make([]byte, 0)
	for i := 1; i <= int(count); i++ {
		section = append(section, byte(i))
	}
	m := c.newMapData(section)
	data := c.newBufData(m)
	return data
}

func (c *codecTestData) newBufData(m map[byte]interface{}) []byte {
	return mapping.NewDataCodec(m)
}

func (c *codecTestData) newMapData(section []byte) map[byte]interface{} {
	data := make(map[byte]interface{})
	for i, v := range section {
		switch i % 7 {
		case 0:
			data[v] = "a"
		case 1:
			data[v] = int32(1)
		case 2:
			data[v] = uint32(1)
		case 3:
			data[v] = int64(1)
		case 4:
			data[v] = uint64(1)
		case 5:
			data[v] = float32(1)
		case 6:
			data[v] = float64(1)
		}
	}
	return data
}