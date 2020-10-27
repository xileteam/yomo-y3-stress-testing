package generator

import (
	"encoding/json"
	"errors"
	"fmt"
)

type jsonTestData struct {
	DefaultCount uint
}

func NewJsonTestData() *jsonTestData {
	return &jsonTestData{DefaultCount: 63}
}

func (c *jsonTestData) GenData() []byte {
	return c.GenDataBy(c.DefaultCount)
}

func (c *jsonTestData) GenDataBy(count uint) []byte {
	if count > c.DefaultCount {
		panic(errors.New("out of bounds"))
	}

	section := make([]int, 0)
	for i := 1; i <= int(count); i++ {
		section = append(section, i)
	}
	m := c.newMapData(section)

	data, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return data
}

func (c *jsonTestData) newMapData(section []int) map[string]interface{} {
	data := make(map[string]interface{})
	for i, v := range section {
		key := fmt.Sprintf("k%d", v)
		switch i % 7 {
		case 0:
			data[key] = "a"
		case 1:
			data[key] = int32(1)
		case 2:
			data[key] = uint32(1)
		case 3:
			data[key] = int64(1)
		case 4:
			data[key] = uint64(1)
		case 5:
			data[key] = float32(1)
		case 6:
			data[key] = float64(1)
		}
	}
	return data
}