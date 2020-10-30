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

// debug:
//func (c *codecTestData) newMapData(section []byte) map[byte]interface{} {
//	data := make(map[byte]interface{})
//	for i, v := range section {
//		switch i % 7 {
//		default:
//			data[v] = "a"
//		}
//	}
//	return data
//}

func (c *codecTestData) newMapData(section []byte) map[byte]interface{} {
	data := make(map[byte]interface{})
	for i, v := range section {
		switch i % 7 {
		case 0:
			data[v] = "a" // 0x61
		case 1:
			data[v] = int32(-65) // []byte{0xFF, 0x3F}
		case 2:
			data[v] = uint32(128) // []byte{0x81, 0x00}
		case 3:
			data[v] = int64(-1) // []byte{0x7F}
		case 4:
			data[v] = uint64(18446744073709551615) // []byte{0x7F}
		case 5:
			data[v] = float32(0.25) // []byte{0x3E, 0x80}
		case 6:
			data[v] = float64(23) // []byte{0x40, 0x37}
		}
	}
	return data
}