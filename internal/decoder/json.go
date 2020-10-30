package decoder

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	TokenEnd = '\n'
)

func TakeValueFromJson(key string, buf []byte) interface{} {
	var sbuf = make([]byte, 0)
	for _, c := range buf {
		if c == TokenEnd {
			var m map[string]interface{}
			if err := json.Unmarshal(sbuf, &m); err != nil {
				panic(errors.New(fmt.Sprintf("data is not json: %s", buf)))
			}

			for k, v := range m {
				if k == key {
					return v
				}
			}
		}
		sbuf = append(sbuf, c)
	}
	return nil
}
