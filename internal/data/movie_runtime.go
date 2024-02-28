package data

import (
	"fmt"
	"strconv"
)

type MovieRuntime int32

func (r MovieRuntime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d minutes", r)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}
