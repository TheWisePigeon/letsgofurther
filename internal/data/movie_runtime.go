package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type MovieRuntime int32

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

func (r MovieRuntime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d minutes", r)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}

func (r *MovieRuntime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}
	i32Value, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	*r = MovieRuntime(i32Value)
	return nil
}
